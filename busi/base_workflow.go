package busi

import (
	"github.com/dtm-labs/dtmgrpc/workflow"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"google.golang.org/grpc"
)

// WorkflowStarup 1
func WorkflowStarup(server *grpc.Server) {
	workflow.InitHTTP(dtmServer, Busi+"/workflow/resume")
	workflow.InitGrpc(dtmutil.DefaultGrpcServer, BusiGrpc, server)
}
