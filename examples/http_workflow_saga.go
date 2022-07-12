package examples

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/workflow"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("http_workflow_simple", func() string {
		wfName := "wf_simple"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqHTTP(data)
			_, err := wf.NewBranch().NewRequest().SetBody(req).Post(busi.Busi + "/TransOut")
			if err != nil {
				return err
			}
			_, err = wf.NewBranch().NewRequest().SetBody(req).Post(busi.Busi + "/TransIn")
			if err != nil {
				return err
			}
			return nil
		})
		logger.FatalIfError(err)

		req := &busi.ReqHTTP{Amount: 30}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmimp.MustMarshal(req))
		logger.Infof("result is: %v", err)
		return gid
	})
	AddCommand("http_workflow_saga", func() string {
		wfName := "wf_saga"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqHTTP(data)
			_, err := wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req).Post(busi.Busi + "/TransOutRevert")
				return err
			}).NewRequest().SetBody(req).Post(busi.Busi + "/TransOut")
			if err != nil {
				return err
			}
			_, err = wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req).Post(busi.Busi + "/TransInRevert")
				return err
			}).NewRequest().SetBody(req).Post(busi.Busi + "/TransIn")
			if err != nil {
				return err
			}
			return nil
		})
		logger.FatalIfError(err)

		req := &busi.ReqHTTP{Amount: 30}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmimp.MustMarshal(req))
		logger.FatalIfError(err)
		return gid
	})

	AddCommand("http_workflow_saga_rollback", func() string {
		wfName := "wf_saga_rollback"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqHTTP(data)
			_, err := wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req).Post(busi.Busi + "/TransOutRevert")
				return err
			}).NewRequest().SetBody(req).Post(busi.Busi + "/TransOut")
			if err != nil {
				return err
			}
			_, err = wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req).Post(busi.Busi + "/TransInRevert")
				return err
			}).NewRequest().SetBody(req).Post(busi.Busi + "/TransIn")
			if err != nil {
				return err
			}
			return nil
		})
		logger.FatalIfError(err)

		req := &busi.ReqHTTP{Amount: 30, TransInResult: dtmcli.ResultFailure}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmimp.MustMarshal(req))
		logger.Infof("result is: %v", err)
		return gid
	})

	AddCommand("http_workflow_saga_barrier", func() string {
		wfName := "wf_saga_barrier"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			req := MustUnmarshalReqHTTP(data)
			_, err := wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req).Post(busi.Busi + "/SagaBTransOutCom")
				return err
			}).NewRequest().SetBody(req).Post(busi.Busi + "/SagaBTransOut")
			if err != nil {
				return err
			}
			_, err = wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req).Post(busi.Busi + "/SagaBTransInCom")
				return err
			}).NewRequest().SetBody(req).Post(busi.Busi + "/SagaBTransIn")
			if err != nil {
				return err
			}
			return nil
		})
		logger.FatalIfError(err)

		req := &busi.ReqHTTP{Amount: 30, TransInResult: dtmcli.ResultFailure}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmimp.MustMarshal(req))
		logger.Infof("result is: %v", err)
		return gid
	})
}
