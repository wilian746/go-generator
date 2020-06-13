#!/bin/sh

SEMVER_UP_TYPE=$1

# semver is an CLI to update your version locally
# for study you can see more in https://semver.org/
# for install you can see more in https://github.com/wilian746/semver-cli
if [ "$SEMVER_UP_TYPE" = "release" ]; then
  semver up release
elif [ "$SEMVER_UP_TYPE" = "minor" ]; then
  semver up minor
elif [ "$SEMVER_UP_TYPE" = "major" ]; then
  semver up minor
fi

ACTUAL_RELEASE=$(semver get release)

if [ "$SEMVER_UP_TYPE" = "rollback" ]; then
    sed -i -e "s/$ACTUAL_RELEASE/{{VERSION_NOT_FOUND}}/g" "./internal/commands/version/version.go"
else
    sed -i -e "s/{{VERSION_NOT_FOUND}}/$ACTUAL_RELEASE/g" "./internal/commands/version/version.go"
fi
