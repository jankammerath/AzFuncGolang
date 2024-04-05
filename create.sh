#!/bin/bash
az deployment group create --resource-group DefaultResourceGroup-DEWC --template-file AzFuncGolang.bicep \
    --parameters location=germanywestcentral \
    --parameters storageAccountName=azfuncgolang \
    --parameters functionAppName=AzFuncGolang \
    --parameters functionAppPlanName=AzFuncGolangPlan