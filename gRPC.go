package main

import (
	context "context"
	"log"
	"net"

	timestamp "github.com/golang/protobuf/ptypes"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type WebServiceImpl struct{}

func StartRPC() {
	grpcServer := grpc.NewServer()
	RegisterWebServiceServer(grpcServer, new(WebServiceImpl))

	listener, err := net.Listen("tcp", ":2333")
	if err != nil {
		log.Fatal("listen err: ", err)
	} else {
		log.Println("RPC已在:2333端口上启动")
	}
	reflection.Register(grpcServer)
	grpcServer.Serve(listener)
}

func (WebServiceImpl) GetLatest(ctx context.Context, args *Room) (*Latest, error) {
	log.Println("查询房间", args.Room)
	miao := GetLatest(args.Room)
	time, _ := timestamp.TimestampProto(miao.Date)
	reply := &Latest{
		Room:      uint32(miao.Room),
		Used:      float32(miao.Used),
		Remaining: float32(miao.Remaining),
		Date:      time,
	}
	return reply, nil
}
