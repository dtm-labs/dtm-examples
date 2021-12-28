/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/dtm/dtmcli"
	"github.com/dtm-labs/dtm/dtmcli/logger"
	"github.com/dtm-labs/dtm/dtmsvr"
	"github.com/dtm-labs/dtm/test/busi"
)

func init() {
	addSample("saga_gorm_barrier", func() string {
		logger.Debugf("a busi transaction begin")
		req := &busi.TransReq{Amount: 30}
		saga := dtmcli.NewSaga(dtmsvr.DefaultHttpServer, dtmcli.MustGenGid(dtmsvr.DefaultHttpServer)).
			Add(busi.Busi+"/SagaBTransOutGorm", busi.Busi+"/SagaBTransOutCompensate", req).
			Add(busi.Busi+"/SagaBTransIn", busi.Busi+"/SagaBTransInCompensate", req)
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})

}
