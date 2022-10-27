# Day 2

Run `make update` after editing/updating the `kustomizations/values.yaml` file to push changes into the cluster.

```
#####################################################################################################################
# Day2 Updates. After Keycloak is up and running. Gather SSO-related information e.g. JWKS, realm, host, CA Cert,
# etc and edit this following section
#####################################################################################################################
#sso:
#  oidc:
#    host: keycloak.bigbang.dev
#    realm: myrealm
#  token_url: "https://{{ .Values.sso.oidc.host }}/auth/realms/{{ .Values.sso.oidc.realm }}/protocol/openid-connect/token"
#  auth_url: "https://{{ .Values.sso.oidc.host }}/auth/realms/{{ .Values.sso.oidc.realm }}/protocol/openid-connect/auth"
#  # CHANGEME: Add a valid JWKS key here
#  jwks: 'CHANGEME-JWKS-KEY'
#  certificate_authority: |
#    -----BEGIN CERTIFICATE-----
#    ...
#    -----END CERTIFICATE-----
#####################################################################################################################

addons:
#####################################################################################################################
# Day 2 Updates. Add chains below. One per application and add the callback_uri, client_id and client_secret values.
#####################################################################################################################
#  authservice:
#    enabled: true
#    values:
#      chains:
#        myapp:
#          match:
#            header: ":authority"
#            prefix: "myapp"
#          callback_uri: https://myapp.bigbang.dev/login/generic_oauth
#          # CHANGEME: Enter a valid client ID e.g. demo-env_a8604cc9-f5e9-4656-802d-d05624370245_bb8-myapp
#          client_id: "CHANGEME_CLIENT_ID"
#          # CHANGEME: Enter a valid client Secret e.g. fsCUSkwy2kaaSlgN4r4LPYOAvHCqzUk5
#          client_secret: "CHANGEME_CLIENT_SECRET"
#####################################################################################################################

```
