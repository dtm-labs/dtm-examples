package examples

import (
	"database/sql"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/workflow"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("http_workflow_xa", func() string {
		wfName := "wf_xa"
		err := workflow.Register(wfName, func(wf *workflow.Workflow, data []byte) error {
			_, err := wf.NewBranch().DoXa(busi.BusiConf, func(db *sql.DB) ([]byte, error) {
				return nil, busi.SagaAdjustBalance(db, busi.TransOutUID, -30, dtmcli.ResultSuccess)
			})
			if err != nil {
				return err
			}
			_, err = wf.NewBranch().DoXa(busi.BusiConf, func(db *sql.DB) ([]byte, error) {
				return nil, busi.SagaAdjustBalance(db, busi.TransInUID, 30, dtmcli.ResultSuccess)
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
