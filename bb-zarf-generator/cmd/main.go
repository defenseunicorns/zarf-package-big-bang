package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	flux "github.com/fluxcd/source-controller/api/v1beta2"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	kubeyaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"
	sigyaml "sigs.k8s.io/yaml"
)

func main() {
	cloneOptions := &git.CloneOptions{
		URL:           "https://repo1.dso.mil/platform-one/big-bang/bigbang.git",
		ReferenceName: plumbing.ReferenceName("refs/tags/1.47.0"),
	}

	if _, err := os.Stat("./tmp/"); err == nil {
		os.RemoveAll("./tmp/")
	}

	_, err := git.PlainClone("./tmp/bigbang/", false, cloneOptions)

	if err != nil {
		panic(err)
	}

	namespace := "bigbang"

	_ = os.Setenv("HELM_NAMESPACE", namespace)

	// Initialize helm SDK
	actionConfig := new(action.Configuration)
	settings := cli.New()

	// Setup K8s connection
	err = actionConfig.Init(settings.RESTClientGetter(), namespace, "", dummyLog)

	if err != nil {
		panic(err)
	}

	client := action.NewInstall(actionConfig)

	client.DryRun = true
	client.Replace = true // Skip the name check
	client.ClientOnly = true
	client.IncludeCRDs = true
	client.ReleaseName = "bigbang"
	client.Namespace = namespace

	chart, err := loader.Load("./tmp/bigbang/chart")

	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile("./values.yaml")

	if err != nil {
		panic(err)
	}

	var values map[string]any

	err = yaml.Unmarshal(file, &values)

	if err != nil {
		panic(err)
	}

	templatedChart, err := client.Run(chart, values)

	if err != nil {
		panic(err)
	}

	yamls, _ := SplitYAML([]byte(templatedChart.Manifest))

	for _, resource := range yamls {
		manifestKind := resource.GetKind()

		if manifestKind == "GitRepository" {
			contents := resource.UnstructuredContent()
			var gitRepository flux.GitRepository

			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(contents, &gitRepository); err != nil {
				panic(err)
			}
			fmt.Printf("%s@%s\n", gitRepository.Spec.URL, gitRepository.Spec.Reference.Tag)
		}
	}
}

func dummyLog(format string, v ...interface{}) {}

// Taken from defenseunicorns/zarf/src/internal/k8s/common.go
//
// SplitYAML splits a YAML file into unstructured objects. Returns list of all unstructured objects
// found in the yaml. If an error occurs, returns objects that have been parsed so far too.
// Source: https://github.com/argoproj/gitops-engine/blob/v0.5.2/pkg/utils/kube/kube.go#L286
func SplitYAML(yamlData []byte) ([]*unstructured.Unstructured, error) {
	var objs []*unstructured.Unstructured
	ymls, err := splitYAMLToString(yamlData)
	if err != nil {
		return nil, err
	}
	for _, yml := range ymls {
		u := &unstructured.Unstructured{}
		if err := sigyaml.Unmarshal([]byte(yml), u); err != nil {
			return objs, fmt.Errorf("failed to unmarshal manifest: %#v", err)
		}
		objs = append(objs, u)
	}
	return objs, nil
}

// Taken from defenseunicorns/zarf/src/internal/k8s/common.go
//
// splitYAMLToString splits a YAML file into strings. Returns list of yamls
// found in the yaml. If an error occurs, returns objects that have been parsed so far too.
// Source: https://github.com/argoproj/gitops-engine/blob/v0.5.2/pkg/utils/kube/kube.go#L304
func splitYAMLToString(yamlData []byte) ([]string, error) {
	// Similar way to what kubectl does
	// https://github.com/kubernetes/cli-runtime/blob/master/pkg/resource/visitor.go#L573-L600
	// Ideally k8s.io/cli-runtime/pkg/resource.Builder should be used instead of this method.
	// E.g. Builder does list unpacking and flattening and this code does not.
	d := kubeyaml.NewYAMLOrJSONDecoder(bytes.NewReader(yamlData), 4096)
	var objs []string
	for {
		ext := runtime.RawExtension{}
		if err := d.Decode(&ext); err != nil {
			if err == io.EOF {
				break
			}
			return objs, fmt.Errorf("failed to unmarshal manifest: %#v", err)
		}
		ext.Raw = bytes.TrimSpace(ext.Raw)
		if len(ext.Raw) == 0 || bytes.Equal(ext.Raw, []byte("null")) {
			continue
		}
		objs = append(objs, string(ext.Raw))
	}
	return objs, nil
}
