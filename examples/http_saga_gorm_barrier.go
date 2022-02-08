/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/logger"
)

func init() {
	AddCommand("http_saga_gorm_barrier", func() string {
		logger.Debugf("a busi transaction begin")
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, dtmcli.MustGenGid(dtmutil.DefaultHTTPServer)).
			Add(busi.Busi+"/SagaBTransOutGorm", busi.Busi+"/SagaBTransOutCom", req).
			Add(busi.Busi+"/SagaBTransIn", busi.Busi+"/SagaBTransInCom", req)
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})

}
