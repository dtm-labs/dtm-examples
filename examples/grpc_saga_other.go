package examples

import (
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmcli/logger"
	dtmgrpc "github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("grpc_saga_hybrid", func() string {
		req := &busi.ReqGrpc{Amount: 30}
		gid := shortuuid.New()
		saga := dtmgrpc.NewSagaGrpc(dtmutil.DefaultGrpcServer, gid).
			Add(busi.BusiGrpc+"/busi.Busi/TransOut", busi.BusiGrpc+"/busi.Busi/TransOutRevert", req)
		saga.Steps = append(saga.Steps, map[string]string{"action": busi.Busi + "/TransIn", "compensate": busi.Busi + "/TransInRevert"})
		saga.BinPayloads = append(saga.BinPayloads, dtmimp.MustMarshal(&busi.ReqGrpc{Amount: 30}))

		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
