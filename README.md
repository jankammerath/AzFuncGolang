# Azure Function in Golang

This boilerplate code is for running an Azure Function in Golang locally on the Mac and deploy it to Azure as an Azure Function.

## No arm64 support

While AWS allows deploying Lambda functions compiled for `arm64` which will use Amazon's Graviton2 cpus, Azure does not support `arm64`. All Azure functions need to be compiled for `x86_64`, or `amd64` as it is defined the `GOARCH` variable. Hence, this boilerplate has two different scripts. The **Azure Function Tools** do not use containers and thus also require the binary to be compiled for Apple Silicon to run locally on Mac.

## Local invokation

The `invoke.sh` script will compile the binary for an Apple Silicon (`arm64` for `darwin`) and execute it using the locally installed Azure Function Tools.

## Deployment to Azure

There's an initial deployment with which the bicep file is deployed. Afterwards the app itself can be deployed to Azure. Whenever bicep file changes, the initial deployment needs to be re-run.

### Initial deployment

The initial deployment with `az-deploy.sh` will create the resource group in the specified location *(within the script)*, then deploy the resources within that group, run the golang compiler, wait 10 seconds for Azure to set up the resources and finally publish the app itself.

### Recurring deployments

The script `deploy.sh` will compile for `amd64` (or `x86_64`) on `linux` to deploy it to Azure. It will not touch any of the resources created on Azure.

### Deletion of the app

Deletion of the app is simply done by deleting the created resource group.

```
az group delete --name AzFuncGolang_RG --yes
```

## Adapter not required

There's no adapter needed to run `go-gin` on Azure Functions. Azure Functions will simply invoke the webserver that Gin provides and redirect any requests to it.