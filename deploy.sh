#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o ./hellofunc .
func azure functionapp publish AzFuncGolang