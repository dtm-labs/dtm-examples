/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package examples

import (
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/dtm-examples/busi"
	"github.com/dtm-labs/dtm-examples/dtmutil"
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("http_saga_redis", func() string {
		busi.SetRedisBothAccount(10000, 10000)
		req := &busi.ReqHTTP{Amount: 30}
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, shortuuid.New()).
			Add(busi.Busi+"/SagaRedisTransOut", busi.Busi+"/SagaRedisTransOutCom", req).
			Add(busi.Busi+"/SagaRedisTransIn", busi.Busi+"/SagaRedisTransInCom", req)
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
	AddCommand("http_saga_redis_rollback", func() string {
		busi.SetRedisBothAccount(10000, 10000)
		saga := dtmcli.NewSaga(dtmutil.DefaultHTTPServer, shortuuid.New()).
			Add(busi.Busi+"/SagaRedisTransIn", busi.Busi+"/SagaRedisTransInCom", &busi.ReqHTTP{Amount: 30}).
			Add(busi.Busi+"/SagaRedisTransOut", busi.Busi+"/SagaRedisTransOutCom", &busi.ReqHTTP{Amount: 30000})
		logger.Debugf("busi trans submit")
		err := saga.Submit()
		logger.FatalIfError(err)
		return saga.Gid
	})
}
