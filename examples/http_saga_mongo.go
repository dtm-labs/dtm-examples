package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/lithammer/shortuuid"
)

func init() {
	AddCommand("http_saga_mongo", func() string {
		busi.SetMongoBothAccount(10000, 10000)
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, shortuuid.New()).
			Add(busi.Busi+"/SagaMongoTransOut", busi.Busi+"/SagaMongoTransOutCom", req).
			Add(busi.Busi+"/SagaMongoTransIn", busi.Busi+"/SagaMongoTransInCom", req)
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
	AddCommand("http_saga_mongo_rollback", func() string {
		busi.SetMongoBothAccount(10000, 10000)
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, shortuuid.New()).
			Add(busi.Busi+"/SagaMongoTransIn", busi.Busi+"/SagaMongoTransInCom", &busi.TransReq{Amount: 30}).
			Add(busi.Busi+"/SagaMongoTransOut", busi.Busi+"/SagaMongoTransOutCom", &busi.TransReq{Amount: 30000})
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
