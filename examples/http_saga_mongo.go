package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
)

func init() {
	AddCommand("http_saga_mongo", func() string {
		busi.SetMongoBothAccount(10000, 10000)
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, dtmcli.MustGenGid(dtmutil.DefaultHTTPServer)).
			Add(busi.Busi+"/SagaMongoTransOut", busi.Busi+"/SagaMongoTransOutCom", req).
			Add(busi.Busi+"/SagaMongoTransIn", busi.Busi+"/SagaMongoTransInCom", req)
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
	AddCommand("http_saga_mongo_rollback", func() string {
		busi.SetMongoBothAccount(10, 10)
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, dtmcli.MustGenGid(dtmutil.DefaultHTTPServer)).
			Add(busi.Busi+"/SagaMongoTransIn", busi.Busi+"/SagaMongoTransInCom", req).
			Add(busi.Busi+"/SagaMongoTransOut", busi.Busi+"/SagaMongoTransOutCom", req)
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
