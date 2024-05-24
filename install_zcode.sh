#!/bin/bash
url="http://127.0.0.1:11434/v1"
model="codegemma"
echo "installing zcode"
go install -a github.com/carsonfeng/ZCode/cmd/zcode@latest
echo "zcode is installed"
echo "setting config openai.base_url to $url"
zcode config set openai.base_url $url
echo "setting config openai.api_key"
zcode config set openai.api_key ollama
echo "setting config openai.model to $model"
zcode config set openai.model $model
echo "setting config openai.timeout to 120s"
zcode config set openai.timeout 120s
echo "success"

