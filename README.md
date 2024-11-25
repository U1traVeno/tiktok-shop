# tiktok-shop
2024秋季字节青训营电商项目

## How to initialize the project

这里已经写好了一个调用了hertz提供的 api.proto 的 auth.proto, 其他的proto文件也可以参考这个文件进行编写. 

接口文档来源于青训营提供的电商项目方案文档, 稍作了修改使之符合 [proto3的语法规范](https://protobuf.dev/programming-guides/proto3/).

首先需要安装 [hz工具](https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/install/)

```shell
go install github.com/cloudwego/hertz/cmd/hz@latest
```

然后执行以下命令, 确认hz工具安装成功

```shell
hz -v

# 输出
hz version v0.9.1
```

执行以下命令, 初始化hertz项目

```shell
hz new -module github.com/U1traVeno/tiktok-shop
```

这会生成一些初始的项目结构, 你可以在这个基础上进行开发. 文件结构大致像 [hz文档中的这样](https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/layout/)

不过因为我们这时还没有根据proto文件生成代码, 没有文档中的model目录

接下来, 我们需要根据proto文件生成代码. 假设你现在编辑过了auth.proto文件, 那么执行以下命令

```shell
hz update -idl idl/auth.proto
```
- 这里不需要`hz update -idl idl/api.proto`(别问我为啥, 官方文档没说要这么干, 我也不知道为啥.

会看到在biz中, handler, model, router目录下都生成了一些auth相关的代码. goland会告诉你哪些是可以编辑的, 哪些不能编辑.

比如model目录下的代码, 会看到文件顶部的注释写了很大的 DO NOT EDIT, goland也会提示你不要编辑这些文件, 因为这些文件是根据proto文件生成的, 你的修改会被覆盖.

又比如router/auth/middleware.go, 点开每个被goland隐藏的方法体, 上面写了`//Your code here`, 这些是你可以编辑的地方.

基本上只要没有DO NOT EDIT的注释, 你都可以编辑.

接下来, 你可以根据自己的需求, 编写业务逻辑, 生成代码, 编写测试, 运行测试, 运行项目.

不过会看到goland把一大堆代码标红, 这是因为还没下载依赖.

```shell
go mod tidy
```

文件名还是会红, 这是表示git还没跟踪这些文件.

这个时候项目已经可以过编译了(其实在hz new的时候就已经可以过编译了), 你可以运行项目了.

```shell
go build

./tiktok-shop
```
会看到输出中, 已经在监听8888端口了.

```shell 
curl http://localhost:8888/ping
{"message":"pong"}
```


