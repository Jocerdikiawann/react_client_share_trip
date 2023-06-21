package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"syscall/js"
	"time"

	"github.com/Jocerdikiawann/react_client_share_trip/service"
	"github.com/Jocerdikiawann/shared_proto_share_trip/route"
	"github.com/tarndt/wasmws"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var client route.RouteClient

func watch() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		dataChan := make(chan *route.LocationResponse)
		defer close(dataChan)
		go func() {
			service.WatchLocation(client, &route.WatchRequest{GoogleId: args[0].String()}, func(data *route.LocationResponse) {
				fmt.Printf("data recieved : %v", data)
				dataChan <- data
			})
		}()
		return <-dataChan
	})
}

func main() {
	ch := make(chan struct{})

	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	const dialTO = time.Second * 30
	dialCtx, dialCancel := context.WithTimeout(appCtx, dialTO)
	defer dialCancel()

	const websocketURL = "ws://localhost:8888/grpc-proxy"
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.DialContext(
		dialCtx,
		"passthrough:///"+websocketURL,
		grpc.WithContextDialer(wasmws.GRPCDialer),
		grpc.WithDisableRetry(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1*1024*1024*1024)),
		grpc.WithTransportCredentials(creds),
	)

	if err != nil {
		log.Fatalf("Could not gRPC dial: %s; Details: %s", websocketURL, err)
	}

	defer conn.Close()

	client = route.NewRouteClient(conn)
	js.Global().Set("watch", watch())
	<-ch
}
