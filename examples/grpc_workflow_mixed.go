package examples

import (
	"database/sql"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/client/workflow"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("grpc_workflow_mixed", func() string {
		wfName := "wf_mixed"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			var req busi.ReqGrpc
			dtmgimp.MustProtoUnmarshal(data, &req)

			_, err := wf.NewBranch().OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := busi.BusiCli.TransOutRevertBSaga(wf.Context, &req)
				return err
			}).Do(func(bb *dtmcli.BranchBarrier) ([]byte, error) {
				return nil, bb.CallWithDB(dbGet().ToSQLDB(), func(tx *sql.Tx) error {
					return busi.SagaAdjustBalance(tx, busi.TransOutUID, int(-req.Amount), "")
				})
			})
			if err != nil {
				return err
			}

			req2 := &busi.ReqHTTP{Amount: int(req.Amount / 2)}
			_, err = wf.NewBranch().OnCommit(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req2).Post(busi.Busi + "/TccBTransInConfirm")
				return err
			}).OnRollback(func(bb *dtmcli.BranchBarrier) error {
				_, err := wf.NewRequest().SetBody(req2).Post(busi.Busi + "/TccBTransInCancel")
				return err
			}).NewRequest().SetBody(req2).Post(busi.Busi + "/TccBTransInTry")
			if err != nil {
				return err
			}
			_, err = wf.NewBranch().DoXa(busi.BusiConf, func(db *sql.DB) ([]byte, error) {
				return nil, busi.SagaAdjustBalance(db, busi.TransInUID, int(req.Amount/2), dtmcli.ResultSuccess)
			})
			return err
		})
		logger.FatalIfError(err)

		req := &busi.ReqHTTP{Amount: 30}
		gid := shortuuid.New()
		err = workflow.Execute(wfName, gid, dtmimp.MustMarshal(req))
		logger.Infof("result is: %v", err)
		return gid
	})
}
