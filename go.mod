module github.com/dtm-labs/dtm-examples

go 1.15

require (
	github.com/dtm-labs/dtmcli v1.13.5
	github.com/dtm-labs/dtmgrpc v1.13.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lib/pq v1.10.3
	github.com/lithammer/shortuuid v3.0.0+incompatible // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	go.mongodb.org/mongo-driver v1.8.3
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.2.1
	gorm.io/gorm v1.22.2
)

// replace github.com/dtm-labs/dtmcli => /Users/wangxi/dtm/dtmcli

// replace github.com/dtm-labs/dtmgrpc => /Users/wangxi/dtm/dtmgrpc
