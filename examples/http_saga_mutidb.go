package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
)

func init() {
	AddCommand("http_saga_multidb", func() string {
		busi.SetMongoBothAccount(10000, 10000)
		busi.SetRedisBothAccount(10000, 10000)
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, dtmcli.MustGenGid(dtmutil.DefaultHTTPServer)).
			Add(busi.Busi+"/SagaMongoTransOut", busi.Busi+"/SagaMongoTransOutCom", &busi.TransReq{Amount: 30}).
			Add(busi.Busi+"/SagaRedisTransOut", busi.Busi+"/SagaRedisTransOutCom", &busi.TransReq{Amount: 20}).
			Add(busi.Busi+"/SagaBTransIn", busi.Busi+"/SagaBTransInCom", &busi.TransReq{Amount: 50})
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
