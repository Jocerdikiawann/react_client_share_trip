package main

import (
	"log"

	"github.com/Jocerdikiawann/react_client_share_trip/service"
	"github.com/Jocerdikiawann/shared_proto_share_trip/route"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err.Error())
	}
	defer conn.Close()

	client := route.NewRouteClient(conn)
	service.WatchLocation(client, &route.WatchRequest{
		GoogleId: "1",
	})
}
