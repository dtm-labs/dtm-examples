# DTM 示例
本项目含大量示例，并且有详细的教程说明如何一步一步上手开发相关的示例

## 运行dtm服务器
您要运行这里的例子，需要首先运行dtm服务器，运行的方式，可以选择最常见的源码运行方式：

``` bash
git clone https://github.com/dtm-labs/dtm && cd dtm
go run main.go
```

成功运行后，dtm会在本地监听两个端口，36789的HTTP，和36790的gRPC

您还可以使用其他多种方式，具体参考[运行dtm服务](https://dtm.pub/)

## 最简quick-start
我们从一个最简单的quick-start开始我们的第一个例子，下面的命令可以运行它

``` bash
git clone https://github.com/dtm-labs/dtm-examples && cd dtm-examples
go run main.go qs

```

运行上述这个例子中，可以看到输出TransOut TransIn，那么一个跨行转账的分布式事务，包含了转出和转入两个子事务，就已经成功完成

## 更多的示例
当您运行完上述例子，有了最初的概念之后，下面我们来详细讲解更多的例子。

examples里面的例子中，quick_start做成了完全独立不共享代码的应用，这样方便使用者快速复制出代码，然后进行修改并上线自己的应用。其他的例子则封装可能会多一些，但原理与quick_start是近似的，用户需要详细研究的话，自己跟踪调试一下代码即可。

例子分为多个维度，包括协议、事务模式、数据库等，你可以运行如下命令，查看所有的示例名称：

``` bash
go run main.go
```

### 协议
dtm支持http协议和gRPC协议，在我们的例子中，文件名里包含了协议名称

### 事务模式
dtm支持四种事务模式：Saga、Tcc、Xa、二阶段消息（Msg），在我们的例子中，文件名包含了事务模式的名称

如果您对这些事务模式还不够熟悉，可以参考[dtm.pub](https://dtm.pub)中的相关讲解

### 数据库
分布式事务通常是要将多个本地事务组合成一个整体全局事务，因此大量的实际业务是带着数据库的。dtm也首创了子事务屏障，帮助用户更好的解决空补偿、悬挂、幂等相关问题。

在带有barrier的示例中，会演示与数据库相关的技巧。运行这部分例子，需要您在如下位置配置好examples连接的数据库：

``` Go
  // main.go -> main()
	busi.BusiConf = dtmimp.DBConf{
		Driver: "mysql",
		Host:   "localhost",
		Port:   3306,
		User:   "root",
	}
```

并且您需要在数据库中创建本示例项目所需要的表，相关的建表sql可以在这里[dtm-sqls](https://github.com/dtm-labs/dtm/tree/main/sqls)找到

在你的数据库中执行这两个sql（假设您的数据库为mysql）

```
dtmcli.barrier.mysql.sql
examples.mysql.sql
```

如果以下命令成功运行，则您的数据库已成功配置

```
go run main.go http_saga_barrier
```

### 第三方ORM
在我们的示例中，还演示了与第三方ORM例如gorm的对接，未来会看需求给出更多第三方ORM的例子

详细对接参考文件：http_xa_gorm http_saga_gorm_barrier


