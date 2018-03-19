#!/usr/bin/env bash

package_name=$1
if [[ -z "$package_name" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

buildFolder="${PWD}/build"

platforms=("linux/amd64" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$buildFolder/$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  
    echo $output_name
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name 
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done