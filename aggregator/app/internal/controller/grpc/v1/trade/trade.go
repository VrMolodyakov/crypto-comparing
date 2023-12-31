package v1

import (
	"fmt"
	"io"

	tradePb "github.com/VrMolodyakov/crypto-comparing/aggregator/app/gen/go/proto/aggregator_service/trade/v1"
)

func (s *server) Create(stream tradePb.TradeService_CreateServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&tradePb.CreateResponse{})
		}
		if err != nil {
			return err
		}
		fmt.Println(req)
	}
}
