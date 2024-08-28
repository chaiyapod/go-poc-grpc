package main

import (
	"log"
	pb "poc/proto"
	"poc/shared"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data []*pb.Data `json:"data"`
}

func getData(ctx *fiber.Ctx) error {
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

	res := Response{
		Data: dataRes,
	}

	return ctx.Status(200).JSON(res)

}
func main() {
	app := fiber.New(fiber.Config{
		AppName: "Poc",
	})

	app.Get("/", getData)
	log.Fatal(app.Listen(":3000"))
}
