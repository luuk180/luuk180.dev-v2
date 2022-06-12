#!/bin/sh
rm main.zip
GOOS=linux GOARCH=amd64 go build -o main maincheck.go
zip main main
rm main