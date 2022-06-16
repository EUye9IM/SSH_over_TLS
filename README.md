# SSH_over_TLS

基于TLS的SSH

## 项目结构

```
ssh_over_tls
 |- client_cli                由 golang 实现的命令行交互客户端
 |   |- config.go             命令行参数解析与配置文件读取
 |   `- main.go               客户端主程序
 |
 |- img                       报告中图片
 |   `- ...
 |
 |- server                    由 golang 实现的服务端
 |   |- config.go             命令行参数解析与配置文件读取
 |   `- main.go               服务端主程序
 |
 |- vendor                    golang 包管理相关文件
 |   `- ...
 |
 |- build-linux-amd64.bat     Windows 下交叉编译 linux 下可执行文件脚本
 |- build-windows-amd64.bat   Windows 下编译 Windows 下可执行文件脚本
 |- conf.yml                  服务端配置文件
 |- go.mod                    golang 包管理相关文件
 |- go.sum                    golang 包管理相关文件
 |- LICENSE                   项目开源许可证
 |- README.md                 自述文件
 |- report.md                 相关课程部分报告
 `- sign.sh                   生成私钥(pem.key)与证书(pem.cer)的脚本

```

## 配置文件说明

### client

参考[tlsssh-client.yml](tlsssh-client.yml)

### server

参考[tlsssh-server.yml](tlsssh-server.yml)

## 可执行文件参数说明

```
tlsssh-client
  -H string
        host (default "localhost")
  -P int
        port (default 443)
  -c string
        path to certificate file (default "pem.cer")
  -d    on debug mode: skipping verify and print code location
  -p string
        password
  -u string
        username.

tlsssh-server -c string
        设定配置文件路径 (默认 "conf.yml")
```

## 环境搭建说明

1. 安装 Golang
2. 运行 build-*.bat 脚本
3. 修改配置文件与 sign.sh 中 ip 地址相关参数
4. 运行 sign.sh 生成密钥与证书
5. 运行服务器
6. 运行客户端

## Q&A

### 我的证书怎么都会连接错误

查看 sign.sh 中 host 的设置，或者尝试在客户机上生成证书。