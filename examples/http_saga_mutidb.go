package examples

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("http_saga_multidb", func() string {
		busi.SetMongoBothAccount(10000, 10000)
		busi.SetRedisBothAccount(10000, 10000)
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, shortuuid.New()).
			Add(busi.Busi+"/SagaMongoTransOut", busi.Busi+"/SagaMongoTransOutCom", &busi.ReqHTTP{Amount: 30}).
			Add(busi.Busi+"/SagaRedisTransOut", busi.Busi+"/SagaRedisTransOutCom", &busi.ReqHTTP{Amount: 20}).
			Add(busi.Busi+"/SagaBTransIn", busi.Busi+"/SagaBTransInCom", &busi.ReqHTTP{Amount: 50})
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
	AddCommand("http_saga_multidb_rollback", func() string {
		busi.SetMongoBothAccount(10000, 10000)
		busi.SetRedisBothAccount(10000, 10000)
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, shortuuid.New()).
			Add(busi.Busi+"/SagaRedisTransOut", busi.Busi+"/SagaRedisTransOutCom", &busi.ReqHTTP{Amount: 20}).
			Add(busi.Busi+"/SagaBTransIn", busi.Busi+"/SagaBTransInCom", &busi.ReqHTTP{Amount: 50}).
			Add(busi.Busi+"/SagaMongoTransOut", busi.Busi+"/SagaMongoTransOutCom", &busi.ReqHTTP{Amount: 30000})
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
