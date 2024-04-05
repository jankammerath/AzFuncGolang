#!/bin/bash
resourceGroupName="AzFuncGolang_RG"
regionName="germanywestcentral"

echo "Creating the resource group..."
az group create --name $resourceGroupName --location $regionName

echo "Deploying the resources..."
az deployment group create --resource-group $resourceGroupName --template-file AzFuncGolang.bicep \
    --parameters location=$regionName \
    --parameters storageAccountName=azfuncgolang \
    --parameters functionAppName=AzFuncGolang \
    --parameters functionAppPlanName=AzFuncGolangPlan

echo "Building the Go function..."
GOOS=linux GOARCH=amd64 go build -o ./hellofunc .

echo "Publishing the function..."
func azure functionapp publish AzFuncGolang