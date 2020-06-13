#!/bin/sh

SEMVER_UP_TYPE=$1

if [[ -z "$SEMVER_UP_TYPE" ]]; then
  echo "SEMVER_UP_TYPE has not been specified, check and try again!"
  exit 1
fi

chmod +x "./deployments/scripts/setup_version.sh"
"./deployments/scripts/setup_version.sh" "$SEMVER_UP_TYPE"

VERSION=$(semver get release)

grep -q "{{VERSION_NOT_FOUND}}" "./internal/commands/version/version.go"
if [ $? -eq 0 ]; then
  echo "The version has not been changed, check and try again!"
  exit 1
fi

git add .
git commit -m "[skip_ci] Change Version"
git tag $VERSION
if [ $? -eq 0 ]; then
  git push origin tag $VERSION
fi

docker build -t wilian746/go-generator:$VERSION -f ./deployments/Dockerfile .
docker push wilian746/go-generator:$VERSION

docker build -t wilian746/go-generator:latest -f ./deployments/Dockerfile .
docker push wilian746/go-generator:latest

"./deployments/scripts/setup_version.sh" "rollback"
git add .
git commit -m "[skip_ci] Change Version"

make build