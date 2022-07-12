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
	"github.com/lithammer/shortuuid/v3"
)

func init() {
	AddCommand("http_msg_barrier_redis", func() string {
		busi.SetRedisBothAccount(10000, 10000)
		gid := shortuuid.New()
		req := busi.GenReqHTTP(30, false, false)
		msg := dtmcli.NewMsg(DtmServer, gid).
			Add(busi.Busi+"/SagaRedisTransIn", req)
		err := msg.DoAndSubmit(busi.Busi+"/RedisQueryPrepared", func(bb *dtmcli.BranchBarrier) error {
			return bb.RedisCheckAdjustAmount(busi.RedisGet(), busi.GetRedisAccountKey(busi.TransOutUID), -30, 86400)
		})
		logger.FatalIfError(err)
		return msg.Gid
	})
	AddCommand("http_msg_barrier_redis_db", func() string {
		busi.SetRedisBothAccount(10000, 10000)
		gid := shortuuid.New()
		req := busi.GenReqHTTP(30, false, false)
		msg := dtmcli.NewMsg(DtmServer, gid).
			Add(busi.Busi+"/SagaBTransIn", req)
		err := msg.DoAndSubmit(busi.Busi+"/RedisQueryPrepared", func(bb *dtmcli.BranchBarrier) error {
			return bb.RedisCheckAdjustAmount(busi.RedisGet(), busi.GetRedisAccountKey(busi.TransOutUID), -30, 86400)
		})
		logger.FatalIfError(err)
		return msg.Gid
	})
}
