package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmgrpc/dtmgimp"
)

func MustUnmarshalReqGrpc(data []byte) *busi.ReqGrpc {
	var req busi.ReqGrpc
	dtmgimp.MustProtoUnmarshal(data, &req)
	return &req
}

func MustUnmarshalReqHTTP(data []byte) *busi.ReqHTTP {
	var req busi.ReqHTTP
	dtmimp.MustUnmarshal(data, &req)
	return &req
}
