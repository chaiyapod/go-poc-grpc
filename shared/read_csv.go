package shared

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Sku struct {
	Guid                     string `csv:"guid"`
	ProductGuid              string `csv:"product_guid"`
	CategoryGuid             string `csv:"category_guid"`
	SellerId                 string `csv:"seller_id"`
	SellerSkuId              string `csv:"seller_sku_id"`
	Code                     string `csv:"code"`
	Slug                     string `csv:"slug"`
	Name                     string `csv:"name"`
	Description              string `csv:"description"`
	SellingSize              string `csv:"selling_size"`
	SellingUnit              string `csv:"selling_unit"`
	ConversionVolume         string `csv:"conversion_volume"`
	ConversionUnit           string `csv:"conversion_unit"`
	ContainAmount            string `csv:"contain_amount"`
	ContainAmountUnit        string `csv:"contain_amount_unit"`
	SeqNo                    string `csv:"seq_no"`
	VisibilityScope          string `csv:"visibility_scope"`
	CreatedAt                string `csv:"created_at"`
	CreatedBy                string `csv:"created_by"`
	UpdatedAt                string `csv:"updated_at"`
	UpdatedBy                string `csv:"updated_by"`
	DeletedAt                string `csv:"deleted_at"`
	DeletedBy                string `csv:"deleted_by"`
	DeliveryTimeSlotEnd      string `csv:"delivery_time_slot_end"`
	DeliveryTimeSlotStart    string `csv:"delivery_time_slot_start"`
	PromoCodeApplicableLevel string `csv:"promo_code_applicable_level"`
	IsSplitOrder             string `csv:"is_split_order"`
	DeliveryDateEnd          string `csv:"delivery_date_end"`
	DeliveryDateStart        string `csv:"delivery_date_start"`
	TradingType              string `csv:"trading_type"`
	PurchaseType             string `csv:"purchase_type"`
	DeliveryTemperature      string `csv:"delivery_temperature"`
	RemainingAmount          string `csv:"remaining_amount"`
	QuotaLimitStatus         string `csv:"quota_limit_status"`
	DeliveryTimePlusDay      string `csv:"delivery_time_plus_day"`
	IsSellByWeight           string `csv:"is_sell_by_weight"`
}

func ReadCsvData() []Sku {
	file, err := os.Open("./__mock__/skus.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var skus []Sku

	err = gocsv.Unmarshal(file, &skus)
	if err != nil {
		fmt.Println("\nAn error occurred while parsing the CSV file " + err.Error())
		return nil
	}

	return skus
}
