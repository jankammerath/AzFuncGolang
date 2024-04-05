# Azure Function in Golang

This boilerplate code is for running an Azure Function in Golang locally on the Mac and deploy it to Azure as an Azure Function.

## No arm64 support

While AWS allows deploying Lambda functions compiled for `arm64` which will use Amazon's Graviton2 cpus, Azure does not support `arm64`. All Azure functions need to be compiled for `x86_64` or `amd64` as it is defined the GOARCH. Hence, this boilerplate has two different scripts. The **Azure Function Tools** do not use containers and hence require the binary to be compiled for Apple Silicon to run locally on Mac.

## Local invokation

The `invoke.sh` script will compile the binary for an Apple Silicon (`arm64` for `darwin`) and execute it using the locally installed Azure Function Tools.

## Deployment to Azure

The script `deploy.sh` will compile for `amd64` (or `x86_64`) on `linux` to deploy it to Azure. 

## Adapter not required

There's no adapter needed to run `go-gin` on Azure Functions. 