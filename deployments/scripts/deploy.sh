#!/bin/sh

VERSION=$1

if [ "$VERSION" = "" ]; then
  VERSION="latest"
fi

chmod +x "./deployments/scripts/setup_version.sh"
"./deployments/scripts/setup_version.sh"

grep -q "{{VERSION_NOT_FOUND}}" "./internal/commands/version/version.go"
if [ $? -eq 0 ]; then
  echo "The version has not been changed, check and try again!"
  exit 1
fi

git tag $VERSION
if [ $? -eq 0 ]; then
  git push origin tag $VERSION
fi

docker build -t wilian746/go-generator:$VERSION -f ./deployments/Dockerfile .
docker push wilian746/go-generator:$VERSION

sh "./deployments/scripts/setup_version.sh" "rollback"
