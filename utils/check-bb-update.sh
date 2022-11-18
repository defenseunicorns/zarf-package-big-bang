#!/bin/bash
#
# This script exits with a 0 if there is an update available. 1 otherwise.
#
NEW_VERSION=$(curl -s https://repo1.dso.mil/api/v4/projects/2872/releases/ | jq '.[0]' | jq -r '.tag_name')
CURRENT_VERSION=$(cat VERSION)

# Versions are in the format: MAJOR.MINOR.PATCHLEVEL. We need to check each level.
#
NV_MAJOR=$(echo $NEW_VERSION | awk -F. '{print $1}')
NV_MINOR=$(echo $NEW_VERSION | awk -F. '{print $2}')
NV_PATCHLEVEL=$(echo $NEW_VERSION | awk -F. '{print $3}')

CV_MAJOR=$(echo $CURRENT_VERSION | awk -F. '{print $1}')
CV_MINOR=$(echo $CURRENT_VERSION | awk -F. '{print $2}')
CV_PATCHLEVEL=$(echo $CURRENT_VERSION | awk -F. '{print $3}')

# Ignore all the release candidate versions
#
echo $NV_PATCHLEVEL | grep rc
if [[ $? -eq 0 ]]; then
	exit 1
fi

if [[ $NV_MAJOR -gt $CV_MAJOR ]]; then
	exit 0
elif [[ $NV_MAJOR -eq $CV_MAJOR ]]; then
	if [[ $NV_MINOR -gt $CV_MINOR ]]; then
		exit 0
	elif [[ $NV_MINOR -eq $CV_MINOR ]]; then
		if [[ $NV_PATCHLEVEL -gt $CV_PATCHLEVEL ]]; then
			exit 0
		fi
	fi
fi
exit 1
