package service

import (
	"context"
	"io"
	"log"

	"github.com/Jocerdikiawann/shared_proto_share_trip/route"
)

func WatchLocation(client route.RouteClient, request *route.WatchRequest, onData func(*route.LocationResponse)) {
	response, err := client.WatchLocation(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}
	ctx := response.Context()
	dataChan := make(chan *route.LocationResponse)
	go func() {
		for {
			resp, err := response.Recv()
			if err == io.EOF {
				close(dataChan)
				return
			}
			if err != nil {
				log.Fatalf("cannot recieve %v", err.Error())
			}
			dataChan <- resp
		}
	}()

	for data := range dataChan {
		onData(data)
	}

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(dataChan)
	}()
}
