package sqlexec

import (
	"database/sql"
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ipfs/go-cid"
	"log"
	"time"
)

type Cluster struct {
	Name  string
	Miner string
}

func InitDB(DSN string) (DB *sql.DB, err error) {

	DB, err = sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}

	info := fmt.Sprintf("dsn check success")
	log.Println(info)

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	info = fmt.Sprintf("database connect success")
	log.Println(info)

	return DB, nil
}

func GetCluster(db *sql.DB) (clusterList map[string]string, err error) {
	clusterList = make(map[string]string)

	sql := "SELECT name,f0 FROM cluster_list"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Name, Miner string

	for rows.Next() {
		err = rows.Scan(&Name, &Miner)
		if err != nil {
			return nil, err
		}
		clusterList[Miner] = Name
	}
	return clusterList, err
}

func Record(db *sql.DB, height abi.ChainEpoch, t time.Time, baseFee int64, cluster string, miner string, cid cid.Cid, methodNum abi.MethodNum, method string, exit exitcode.ExitCode, value abi.TokenAmount,cost abi.TokenAmount) (SQL string,err error) {
	SQL = fmt.Sprintf(
		"INSERT INTO filecoin_cluster_message_record" +
		"(height,time,basefee,cluster,miner,cid,method_num,method,exit_code,`exit`,value,cost)"+
		"VALUES('%d','%s','%d','%s','%s','%s','%d','%s','%d','%s','%d','%d')",
		height,
		t.Format("2006-01-02 15:04:05"),
		baseFee,
		cluster,
		miner,
		cid,
		methodNum,
		method,
		exit,
		exit,
		value,
		cost,
	)

	_, err = db.Exec(SQL)
	return SQL, err
}
