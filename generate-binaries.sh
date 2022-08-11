#!/usr/bin/env bash
set -e

# echo "Running tests before generating the binaries..."
# go test -v -cover
echo "Proceeding with generating the binaries..."

platforms=("linux/amd64" "darwin/amd64" "darwin/arm64")
VERSION=`grep -E "MAJOR|MINOR|PATCH" cmd/revealer/version.go | cut -d '"' -f 2 | xargs echo -n | tr -s " " "."`

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=revealer'-v'$VERSION'-'$GOOS'-'$GOARCH

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $PWD/cmd/revealer
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
