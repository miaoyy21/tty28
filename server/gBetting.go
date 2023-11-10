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

type BettingService struct {
	pb.UnimplementedBettingServiceServer
}
type BettingResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (s *BettingService) Betting(ctx context.Context, r *pb.BettingRequest) (*pb.BettingResponse, error) {
	var resp BettingResponse

	err := hdo.Do(r.GetAuthority(), r.GetOrigin(), r.GetReferer(), r.GetSecChUa(), r.GetSecChUaPlatform(), r.GetUserAgent(), r.GetUrl(), &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Code, resp.Msg)
	}

	log.Println("执行投注成功 ...")
	return &pb.BettingResponse{}, nil
}

func gBetting(port string) {
	server := grpc.NewServer()
	pb.RegisterBettingServiceServer(server, &BettingService{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Gold net.Listen() Failure : %s \n", err.Error())
	}

	if err := server.Serve(l); err != nil {
		log.Fatalf("Gold server.Serve() Failure : %s \n", err.Error())
	}
}
