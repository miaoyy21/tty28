package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"tty28/hdo"
	pb "tty28/proto"
)

type GoldService struct {
	pb.UnimplementedGoldServiceServer
}

type QGoldResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Gold string `json:"gold"`
	} `json:"data"`
}

func (s *GoldService) Gold(ctx context.Context, r *pb.GoldRequest) (*pb.GoldResponse, error) {
	var resp QGoldResponse

	err := hdo.Do(r.GetAuthority(), r.GetOrigin(), r.GetReferer(), r.GetSecChUa(), r.GetSecChUaPlatform(), r.GetUserAgent(), r.GetUrl(), &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Code, resp.Msg)
	}

	gold, err := hdo.Int64(resp.Data.Gold)
	if err != nil {
		return nil, err
	}

	return &pb.GoldResponse{Gold: gold}, nil
}

func gGold(port string) {
	server := grpc.NewServer()
	pb.RegisterGoldServiceServer(server, &GoldService{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Gold net.Listen() Failure : %s \n", err.Error())
	}

	if err := server.Serve(l); err != nil {
		log.Fatalf("Gold server.Serve() Failure : %s \n", err.Error())
	}
}
