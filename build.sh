#!/bin/bash

distPath="./bin"
mkdir -p ${distPath}

platforms=("linux/amd64" "linux/arm64" "windows/amd64" "darwin/amd64" "darwin/arm64")
export CGO_ENABLED=0

for platform in "${platforms[@]}"; do
    operationalSystem=${platform%/*}
    isa=${platform#*/}

    binaryName="poc-${operationalSystem}-${isa}"
    if [ "${operationalSystem}" == "windows" ]; then
        binaryName+=".exe"
    fi

    echo "Building for: ${operationalSystem}/${isa}"

    if ! GOOS="${operationalSystem}" GOARCH="${isa}" go build -o "${distPath}/${binaryName}" ./main.go; then
        echo "An error has occurred! Aborting the script execution..."
        exit 1
    fi
done

echo "Building completed!"
