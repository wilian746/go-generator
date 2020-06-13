#!/bin/sh

VERSION=$1

if [ "$VERSION" = "" ]; then
  VERSION="latest"
fi

chmod +x "./deployments/scripts/setup_version.sh"
sh "./deployments/scripts/setup_version.sh"

git tag $VERSION
if [ $? -eq 0 ]; then
  git push origin tag $VERSION
fi

docker build -t wilian746/go-generator:$VERSION -f ./deployments/Dockerfile .
docker push wilian746/go-generator:$VERSION

sh "./deployments/scripts/setup_version.sh" "rollback"
