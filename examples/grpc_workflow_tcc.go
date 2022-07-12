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
	AddCommand("grpc_workflow_tcc", func() string {
		wfName := "wf_tcc"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqGrpc(data)
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutRevert(wf.Context, req)
				return err
			}).OnCommit(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutConfirm(wf.Context, req)
				return err
			})
			_, err := busi.BusiCli.TransOut(wf.Context, req)
			if err != nil {
				return err
			}
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransInRevert(wf.Context, req)
				return err
			}).OnCommit(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransInConfirm(wf.Context, req)
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
	AddCommand("grpc_workflow_tcc_rollback", func() string {
		wfName := "wf_tcc_rollback"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqGrpc(data)
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutRevert(wf.Context, req)
				return err
			}).OnCommit(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutConfirm(wf.Context, req)
				return err
			})
			_, err := busi.BusiCli.TransOut(wf.Context, req)
			if err != nil {
				return err
			}
			wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransInRevert(wf.Context, req)
				return err
			}).OnCommit(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransInConfirm(wf.Context, req)
				return err
			})
			_, err = busi.BusiCli.TransIn(wf.Context, req)
			return err
		})
		logger.FatalIfError(err)

		req := &busi.ReqGrpc{Amount: 30, TransInResult: "FAILURE"}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmgimp.MustProtoMarshal(req))
		logger.Infof("the result is: %v", err)
		return gid
	})
}
