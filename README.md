English | [简体中文](https://github.com/dtm-labs/dtm-examples/blob/main/doc/README-cn.md)

# DTM Examples
dtm has many examples to help you get started with distributed transactions quickly
- [quick-start-sample](https://github.com/dtm-labs/quick-start-sample): some minimal example of using dtm HTTP/gRPC/workflow client
- [dtm-examples](https://github.com/dtm-labs/dtm-examples): a project contains a lot of examples, mainly demonstrating various uses of the dtm SDK.
- [dtm-cases](https://github.com/dtm-labs/dtm-cases): a project contains several projects, mainly demonstrating the complete application of dtm in some fields, such as order system, flash-sales system.
- [dtmdriver-clients](https://github.com/dtm-labs/dtmdriver-clients): contains dtm support for microservices framework, e.g. go-zero example

## Run the dtm server
If you want to run the example here, you need to run the dtm server first. To run it, you can choose the most common way of running source code: the

``` bash
git clone https://github.com/dtm-labs/dtm && cd dtm
go run main.go
```

After running successfully, dtm will listen locally on two ports, 36789 for HTTP, and 36790 for gRPC.

You can also use a variety of other methods, see [running dtm service](https://en.dtm.pub/)

## The simplest quick-start
Let's start our first example with a minimal quick-start, which can be run with the following command

``` bash
git clone https://github.com/dtm-labs/dtm-examples && cd dtm-examples
go run main.go qs

```

Running this example above, you can see the output TransOut TransIn. A distributed transaction for a cross-row transfer, containing two sub-transactions, TransOut and TransIn, has completed successfully

## Examples of common dimensions
Once you have run the above example and have the first impression, let's explain more examples in detail.

In the examples, quick_start is made to be a completely standalone, non-shared application, so that users can quickly copy out the code, then modify it and run their own application. Other examples may be more encapsulated, but the principle is similar to that of quick_start. Users can trace and debug the code themselves if they need to study it in detail.

The examples are divided into multiple dimensions, including protocols, transaction mode, database, etc. You can run the following command to see all the example names.

``` bash
go run main.go
```

### Protocol classification
dtm supports http protocol and gRPC protocol
- HTTP protocol: dtm server will listen to HTTP port 36789, here the example business will listen to HTTP 8081
- gRPC protocol: dtm server will listen to gRPC port 36790, here the example business will listen to gRPC 58081

In all examples, the name of the protocol will be given, http or grpc. You can select the appropriate one according to your needs.

### Transaction modes
dtm supports multiple transaction modes. In all examples, the names of the transaction modes are included, as follows, you can choose according to your needs
- msg: two-stage message, suitable for global transactions that do not require rollback
- saga: suitable for global transactions that need to support rollback
- tcc: for global transactions with high consistency requirements
- xa: for global transactions with low performance requirements and no rowlock contention

If you are not familiar with these transaction solutions, you can refer to the related explanation in [en.dtm.pub](https://en.dtm.pub)

### Database
Distributed transactions are usually about combining multiple local transactions into one overall global transaction, so a lot of real business is carried out with databases. dtm also pioneered sub-transaction barriers to help users better solve problems related to null compensation, suspensions, powers, etc.

In the example with barrier, database related tricks are demonstrated. dtm provides a sample mysql for your use to save you the trouble of configuring the database.

> If you want to configure your own database, see Deployment and Operations in the [dtm documentation](https://en.dtm.pub)

### Redis
dtm can combine Redis to form a global transaction, just look for examples with redis in the name

### Mongo
dtm can combine Mongo to form a global transaction, just look for examples with mongo in the name

### Mutiple Databases
dtm can combine Mysql, Redis, Mongo to form a global transaction, just look for examples with `multidb` in the name

### Transaction rollback
We have a number of examples that demonstrate rollbacks, look for examples with rollback in their name.

## More features

### Third party ORM
In our examples, we also demonstrate interfacing with third-party ORMs such as gorm, and will give more examples of third-party ORMs as needed

You can look for examples with gorm in their name

### Custom header
Some subtransactions in your business require custom header. dtm supports global transaction granularity header customization, i.e. you can specify custom header for each global transaction, and dtm will add your specified header when it calls your subtransaction service.

HTTP and gRPC both support custom header, see the example with Header in the name for details
