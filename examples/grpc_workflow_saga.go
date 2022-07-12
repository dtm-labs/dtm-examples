package examples

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/client/workflow"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("grpc_workflow_simple", func() string {
		wfName := "wf_simple"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqGrpc(data)
			_, err := busi.BusiCli.TransOut(wf.NewBranchCtx(), req)
			if err != nil {
				return err
			}
			_, err = busi.BusiCli.TransIn(wf.NewBranchCtx(), req)
			return err
		})
		logger.FatalIfError(err)

		req := &busi.ReqGrpc{Amount: 30}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmgimp.MustProtoMarshal(req))
		logger.FatalIfError(err)
		return gid
	})
	AddCommand("grpc_workflow_saga", func() string {
		wfName := "wf_saga"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqGrpc(data)
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutRevert(wf.Context, req)
				return err
			})
			_, err := busi.BusiCli.TransOut(wf.Context, req)
			if err != nil {
				return err
			}
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransInRevert(wf.Context, req)
				return err
			})
			_, err = busi.BusiCli.TransIn(wf.Context, req)
			return err
		})
		logger.FatalIfError(err)

		req := &busi.ReqGrpc{Amount: 30}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmgimp.MustProtoMarshal(req))
		logger.FatalIfError(err)
		return gid
	})
	AddCommand("grpc_workflow_saga_rollback", func() string {
		wfName := "wf_saga_rollback"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqGrpc(data)
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutRevert(wf.Context, req)
				return err
			})
			_, err := busi.BusiCli.TransOut(wf.Context, req)
			if err != nil {
				return err
			}
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransInRevert(wf.Context, req)
				return err
			})
			_, err = busi.BusiCli.TransIn(wf.Context, req)
			return err
		})
		logger.FatalIfError(err)

		req := &busi.ReqGrpc{Amount: 30, TransInResult: dtmcli.ResultFailure}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmgimp.MustProtoMarshal(req))
		logger.Debugf("result is: %v", err)
		return gid
	})
	AddCommand("grpc_workflow_saga_barrier", func() string {
		wfName := "wf_saga_barrier"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqGrpc(data)
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutRevertBSaga(wf.Context, req)
				return err
			})
			_, err := busi.BusiCli.TransOutBSaga(wf.Context, req)
			if err != nil {
				return err
			}
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransInRevertBSaga(wf.Context, req)
				return err
			})
			_, err = busi.BusiCli.TransInBSaga(wf.Context, req)
			return err
		})
		logger.FatalIfError(err)

		req := &busi.ReqGrpc{Amount: 30}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmgimp.MustProtoMarshal(req))
		logger.FatalIfError(err)
		return gid
	})
}
