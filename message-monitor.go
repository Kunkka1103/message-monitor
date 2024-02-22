package main

import (
	"context"
	"flag"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v0api"
	"log"
	"message-monitor/analyzer"
	"message-monitor/prometh"
	"message-monitor/sqlexec"
	"net/http"
	"time"
)

var lotusAPI = flag.String("l", "http://127.0.0.1:1234/rpc/v0", "lotusAPI")
var dsn = flag.String("d", "", "ops dsn")
var pushGatewayAddr = flag.String("p", "", "pushgateway addr")

func ConnectClient(apiUrl string) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	header := http.Header{}
	ctx := context.Background()
	return client.NewFullNodeRPCV0(ctx, apiUrl, header)
}

func main() {

	flag.Parse()

	if *dsn==""{
		log.Fatalf("ops dsn must be provided")
	}

	//init lotus connect
	delegate, closer, err := ConnectClient(*lotusAPI)
	if err != nil {
		log.Fatalf("connect to lotusAPI failed,err:%s", err)
	}
	defer closer()

	//init db
	db, err := sqlexec.InitDB(*dsn)
	if err != nil {
		log.Fatalf("db init failed,err:%s", err)
	}

	//gen ctx
	ctx := context.Background()

	//get base
	startTS, err := delegate.ChainHead(ctx)
	if err != nil {
		log.Printf("curTSK get failed,err:%s", err)
	}
	baseHeight := startTS.Height()
	log.Printf("curTsk get success,latest height:%d", baseHeight)

	for {
		latestTS, err := delegate.ChainHead(ctx)
		if err != nil {
			log.Printf("latestTS get failed,err:%s", err)
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("latestTS get success")

		latestHeight := latestTS.Height()
		log.Printf("latest height get success:%d", latestHeight)

		//compare base and latest
		if latestHeight == baseHeight {
			log.Printf("latest height is equal base height,pass")
			time.Sleep(10 * time.Second)
			continue
		}

		baseFee := latestTS.Blocks()[0].ParentBaseFee.Int64()

		clusterList, err := sqlexec.GetCluster(db)
		if err != nil {
			log.Printf("get clusterlist failed,err:%s", err)
			time.Sleep(10 * time.Second)
			continue
		}

		log.Printf("Parsing messages at %v", latestHeight)

		Messages, err := delegate.ChainGetMessagesInTipset(ctx, latestTS.Parents())
		if err != nil {
			log.Printf("message parse failed at epoch %d", latestHeight-1)
			continue
		}

		prometh.Push(float64(len(Messages)), *pushGatewayAddr)


		for _, m := range Messages {
			analyzer.Analyzer(db, latestHeight-1,baseFee, clusterList, m, delegate, ctx)
		}

		baseHeight = latestHeight
		time.Sleep(10 * time.Second)

	}

}
