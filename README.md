# SSH_over_TLS
基于TLS的SSH

## 项目结构

```
ssh_over_tls
 |
 |- client_cli         由 golang 实现的命令行交互客户端
 |   |- conf.yml       客户端配置文件
 |   |- config.go      命令行参数解析与配置文件读取
 |   `- main.go        客户端主程序
 |
 |- server             由 golang 实现的服务端
 |   |- conf.yml       服务端配置文件
 |   |- config.go      命令行参数解析与配置文件读取
 |   `- main.go        服务端主程序
 |
 |- sign               签名存放文件夹
 |   |- server.csr     服务端证书
 |   |- server.key     服务端私钥
 |   `- sign.sh        生成私钥与证书的脚本
 |
 |- vendor             golang 包管理相关文件
 |
 |- build.bat          编译可执行文件脚本
 |- go.mod             golang 包管理相关文件
 |- go.sum             golang 包管理相关文件
 |- LICENSE            项目开源许可证
 `- README.md          自述文件

```

## 配置文件说明

### client-cli

参考[client_cli/conf.yml](client_cli/conf.yml)

### server

参考[server/conf.yml](server/conf.yml)

## 可执行文件参数说明

```
client_cli.exe -c string
        设定配置文件路径 (默认 "client_cli/conf.yml")

server.exe -c string
        设定配置文件路径 (默认 "server/conf.yml")
```

## 环境搭建说明

1. 安装 Golang

	（略）

2. 运行 build.bat 脚本

	（略）

3. 打开服务器

	双击 server 可执行文件

4. 打开客户端

	双击 client_cli 可执行文件