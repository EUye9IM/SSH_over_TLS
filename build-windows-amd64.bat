@echo off
go build -o tlsssh-server.exe -mod=vendor ./server
go build -o tlsssh-client.exe -mod=vendor ./client