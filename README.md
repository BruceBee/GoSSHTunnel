## GoSSHTunnel使用手册

> GoSSHTunnel是一个用Golang开发的SSH转发小脚本，主要用于连接远程主机上的服务。即远程服务拒绝公网直接连接的情况下，通过SSH通道进行连接的一种方式。
>
> 版本 v1.0

### 1、私钥替换

将`privateKey`目录下的id_rsa替换成自己的私钥文件


### 2、配置修改

修改`conf/app.toml`

```shell script

ssh_host = "xxx.xxx.xxx" # 登录IP
ssh_port = 22 # 登录端口
ssh_user = "root" # 登录账号
ssh_pkey = "./privateKey/id_rsa"  # 确保秘钥文件路径正确
ssh_pass = "xxxx" # 修改成自己的秘钥密码，没有则为空

# 绑定端口(建立SSH通道后,指向数据库端口, port1:port2,其中port1代表本地端口，port2代表远程端口)
bind_port = [
            [12345, 12345],
            [22345, 22345],
            [33345, 33345]
]
```

###  3、执行即可

编译

```
> go build
```




Windows:
```shell script
> cd $PATH/GoSSHTunnel
> GoSSHTunnel.exe
```

Linux/MacOS:

```shell script

$ cd $PATH/GoSSHTunnel
$ ./GoSSHTunnel

```

### 4、连接方式

执行脚本以后，本地实际上就和远程主机为“同一台主机”。
因此连接数据的地址，就是`127.0.0.1:端口`

