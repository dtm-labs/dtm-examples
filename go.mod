module github.com/dtm-labs/dtm-examples

go 1.15

require (
	github.com/dtm-labs/client v1.17.3
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lib/pq v1.10.3
	github.com/lithammer/shortuuid/v3 v3.0.7
	go.mongodb.org/mongo-driver v1.9.1
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/postgres v1.2.1
	gorm.io/gorm v1.22.2
)

// replace github.com/dtm-labs/client/dtmcli => /Users/wangxi/dtm/dtmcli

// replace github.com/dtm-labs/client/dtmgrpc => /Users/wangxi/dtm/dtmgrpc
