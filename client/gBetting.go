package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "tty28/proto"
)

func gBetting(target, issue, sBets, uToken, secChUa, secChUaPlatform, userAgent string, ns int) error {
	// Create a client connection to the given target with a credentials which disables transport security
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewBettingServiceClient(conn)

	qUrl := fmt.Sprintf("%s?utoken=%s&cid=%s&bet_data=[%s]&stylePath=%s&t=%d", conf.BettingURL, uToken, issue, sBets, conf.Style, ns)
	req := &pb.BettingRequest{
		Url: qUrl,

		Authority: conf.Authority,
		Origin:    conf.Origin,
		Referer:   conf.Referer,

		SecChUa:         secChUa,
		SecChUaPlatform: secChUaPlatform,
		UserAgent:       userAgent,
	}

	if _, err := client.Betting(context.Background(), req); err != nil {
		return err
	}

	return nil
}
