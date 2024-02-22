package analyzer

import (
	"context"
	"database/sql"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/types"
	"log"
	"message-monitor/height"
	"message-monitor/method"
	"message-monitor/sqlexec"
)

func Analyzer(db *sql.DB, h abi.ChainEpoch,baseFee int64, clusterList map[string]string, m api.Message, delegate v0api.FullNode, ctx context.Context) {
	if clusterList[m.Message.To.String()] == "" {
		return
	}

	in, err := delegate.StateReplay(ctx, types.EmptyTSK, m.Cid)
	if err != nil {
		return
	}

	sql,err := sqlexec.Record(
		db,
		h,
		height.ToTime(h),
		baseFee,
		clusterList[m.Message.To.String()],
		m.Message.To.String(),
		m.Cid,
		m.Message.Method,
		method.MethodsMap[m.Message.Method],
		in.MsgRct.ExitCode,
		m.Message.Value,
		in.GasCost.TotalCost,
	)

	if err != nil {
		log.Printf("'%s' exce failed, err:%s", sql, err)
		return
	}

	log.Printf("'%s' exec success", sql)
	return

}
