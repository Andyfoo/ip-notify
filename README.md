# ip-notify
##功能
定时获取外网IP变动并发送到自定邮箱。自家树莓派上的小工具

## 开始

```shell
$dep ensure
```



## 构建

``` shell
$ docker build -t ip-notify:0.1 .  
```



## 部署

```shell
$ docker run -d -it   ip-notify:0.1 
```