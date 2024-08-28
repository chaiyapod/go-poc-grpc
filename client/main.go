package main

import (
	"context"
	"log"
	pb "poc/proto"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Response struct {
	Data []*pb.Data `json:"data"`
}

func getByHttp(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON("res")
}

func getByGrpc(fiberCtx *fiber.Ctx) error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(10*1024*1024), // 10 MB
		grpc.MaxCallSendMsgSize(10*1024*1024), // 10 MB
	))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetSkuData(ctx, &pb.HelloRequest{})
	if err != nil {
		log.Fatalf("could not get sku data: %v", err)
	}

	resData := res.Data

	return fiberCtx.Status(200).JSON(resData)
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Poc",
	})

	app.Get("/http", getByHttp)
	app.Get("/grpc", getByGrpc)
	log.Fatal(app.Listen(":4000"))
}
