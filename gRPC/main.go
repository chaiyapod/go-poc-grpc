package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "poc/proto"
	"poc/shared"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) GetData(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("RequestData: name 1 -> %v, name 2 -> %v", in.GetName1(), in.GetName2())

	skus := shared.ReadCsvData()

	dataRes := []*pb.Data{}

	for _, sku := range skus {
		dataRes = append(dataRes, &pb.Data{
			Guid:                     sku.Guid,
			ProductGuid:              sku.ProductGuid,
			CategoryGuid:             sku.CategoryGuid,
			SellerId:                 sku.SellerId,
			SellerSkuId:              sku.SellerSkuId,
			Code:                     sku.Code,
			Slug:                     sku.Slug,
			Name:                     sku.Name,
			Description:              sku.Description,
			SellingSize:              sku.SellingSize,
			SellingUnit:              sku.SellingUnit,
			ConversionVolume:         sku.ConversionVolume,
			ConversionUnit:           sku.ConversionUnit,
			ContainAmount:            sku.ContainAmount,
			ContainAmountUnit:        sku.ContainAmountUnit,
			SeqNo:                    sku.SeqNo,
			VisibilityScope:          sku.VisibilityScope,
			CreatedAt:                sku.CreatedAt,
			CreatedBy:                sku.CreatedBy,
			UpdatedAt:                sku.UpdatedAt,
			UpdatedBy:                sku.UpdatedBy,
			DeletedAt:                sku.DeletedAt,
			DeletedBy:                sku.DeletedBy,
			DeliveryTimeSlotEnd:      sku.DeliveryTimeSlotEnd,
			DeliveryTimeSlotStart:    sku.DeliveryTimeSlotStart,
			PromoCodeApplicableLevel: sku.PromoCodeApplicableLevel,
			IsSplitOrder:             sku.IsSplitOrder,
			DeliveryDateEnd:          sku.DeliveryDateEnd,
			DeliveryDateStart:        sku.DeliveryDateStart,
			TradingType:              sku.TradingType,
			PurchaseType:             sku.PurchaseType,
			DeliveryTemperature:      sku.DeliveryTemperature,
			RemainingAmount:          sku.RemainingAmount,
			QuotaLimitStatus:         sku.QuotaLimitStatus,
			DeliveryTimePlusDay:      sku.DeliveryTimePlusDay,
			IsSellByWeight:           sku.IsSellByWeight,
		})

	}

	return &pb.HelloResponse{Data: dataRes}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
