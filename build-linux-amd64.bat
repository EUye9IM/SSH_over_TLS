@echo off
SET GOOS=linux
SET GOARCH=amd64
go build -o tlsssh-server -mod=vendor ./server
go build -o tlsssh-client -mod=vendor ./client
