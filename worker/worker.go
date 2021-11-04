package worker

import (
	"time"

	"github.com/acentswap/aswap/v3/router/bridge"
	"github.com/acentswap/aswap/v3/rpc/client"
)

const interval = 10 * time.Millisecond

// StartRouterSwapWork start router swap job
func StartRouterSwapWork(isServer bool) {
	logWorker("worker", "start router swap worker")

	client.InitHTTPClient()
	bridge.InitRouterBridges(isServer)

	bridge.StartAdjustGatewayOrderJob()
	time.Sleep(interval)

	if !isServer {
		StartAcceptSignJob()
		return
	}

	StartSwapJob()
	time.Sleep(interval)

	StartVerifyJob()
	time.Sleep(interval)

	StartStableJob()
	time.Sleep(interval)

	StartReplaceJob()
	time.Sleep(interval)

	StartPassBigValueJob()
}
