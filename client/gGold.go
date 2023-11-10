package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "tty28/proto"
)

func gGold(target, uToken, secChUa, secChUaPlatform, userAgent string, ns int) (int64, error) {
	// Create a client connection to the given target with a credentials which disables transport security
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	client := pb.NewGoldServiceClient(conn)

	req := &pb.GoldRequest{
		Url: fmt.Sprintf("%s?utoken=%s&t=%d", conf.GoldURL, uToken, ns),

		Authority: conf.Authority,
		Origin:    conf.Origin,
		Referer:   conf.Referer,

		SecChUa:         secChUa,
		SecChUaPlatform: secChUaPlatform,
		UserAgent:       userAgent,
	}

	resp, err := client.Gold(context.Background(), req)
	if err != nil {
		return 0, err
	}

	return resp.GetGold(), nil
}
