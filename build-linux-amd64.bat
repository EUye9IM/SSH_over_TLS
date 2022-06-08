@echo off
SET GOOS=linux
SET GOARCH=amd64
go build -o tls-sshd -mod=vendor ./server
go build -o tls-ssh -mod=vendor ./client_cli
