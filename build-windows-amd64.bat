@echo off
go build -o tls-sshd.exe -mod=vendor ./server
go build -o tls-ssh.exe -mod=vendor ./client_cli
