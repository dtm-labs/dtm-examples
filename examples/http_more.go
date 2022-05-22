package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/lithammer/shortuuid"
)

func init() {
	AddCommand("http_saga_multiSource", func() string {
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, shortuuid.New()).
			Add(busi.Busi+"/SagaMultiSource", busi.Busi+"/SagaMultiSourceRevert", req)
		logger.Debugf("saga busi trans submit")
		err := saga.Submit()
		logger.Debugf("result gid is: %s", saga.Gid)
		logger.FatalIfError(err)
		return saga.Gid
	})
}
