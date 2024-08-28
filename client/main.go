package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	pb "poc/proto"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Response struct {
	Data []*pb.Data `json:"data"`
}

var client = resty.New()

func getByHttp(ctx *fiber.Ctx) error {
	start := time.Now()
	resp, err := client.
		NewRequest().
		SetHeader("Content-Type", "application/json").
		Get("https://poc-http-zcjkqsokeq-as.a.run.app")
	elapsed := time.Since(start)

	if err != nil {
		return ctx.Status(400).JSON(err)
	}
	fmt.Println("[HTTP] request took %s", elapsed)

	if resp.StatusCode() != http.StatusOK {
		fmt.Println("error")
		return ctx.Status(resp.StatusCode()).JSON(resp.Error())
	}

	data := Response{}

	if err = json.Unmarshal(resp.Body(), &data); err != nil {
		return ctx.Status(400).JSON(err)
	}

	return ctx.Status(200).JSON(data)
}

func getByGrpc(fiberCtx *fiber.Ctx) error {

	conn, err := grpc.NewClient("poc-grpc-zcjkqsokeq-as.a.run.app", grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(10*1024*1024), // 10 MB
		grpc.MaxCallSendMsgSize(10*1024*1024), // 10 MB
	))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	start := time.Now()

	res, err := c.GetSkuData(ctx, &pb.HelloRequest{})

	elapsed := time.Since(start)
	fmt.Println("[gRPC] request took %s", elapsed)
	if err != nil {
		log.Fatalf("could not get sku data: %v", err)
	}

	resData := res.Data

	return fiberCtx.Status(200).JSON(resData)
}

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Poc",
	})

	app.Get("/http", getByHttp)
	app.Get("/grpc", getByGrpc)
	log.Fatal(app.Listen(":4000"))
}
