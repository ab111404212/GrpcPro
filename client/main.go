package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"

	proto "github.com/ab111404212/pb-proto"
)

const (
	address     = "localhost:50051"
	defaultName = "login"
)

var ipAddress string

// func init() {
// 	flag.StringVar(&ipAddress, "addr", "", "input server address")
// }

func main() {
	flag.StringVar(&ipAddress, "addr", "", "input server address")
	flag.Parse()
	if ipAddress == "" {
		log.Fatalf("failed to get ipAddress: %v", ipAddress)
	} else {
		log.Printf("ipAddress: %v", ipAddress)
	}
	// Set up a connection to the server.
	wg := sync.WaitGroup{}
	wg.Add(1)

	for j := 0; j < 25000; j++ {
		conn, err := grpc.Dial(ipAddress, grpc.WithInsecure())
		if err != nil {
			log.Printf("did not connect: %v", err)
		}
		// defer conn.Close()
		c := proto.NewServerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.LoginRequest(ctx, &proto.UserLoginRequest{Id: int32(j)})
		if err != nil {
			log.Printf("could not Login: %v", err)
		}
		log.Printf("Login:Coin %d", r.GetData().GetCoin())
		time.Sleep(time.Millisecond * 2)
		conn.Close()
	}
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// c := proto.NewServerClient(conn)

	// // Contact the server and print out its response.

	// for i := 0; i < 10; i++ {
	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 	defer cancel()
	// 	r, err := c.LoginRequest(ctx, &proto.UserLoginRequest{Id: int32(i) + 1421412})
	// 	if err != nil {
	// 		log.Fatalf("could not Login: %v", err)
	// 	}
	// 	log.Printf("Login: %s", r.GetData().GetEquips())
	// 	log.Printf("Login: %d", r.GetData().GetCoin())
	// 	log.Printf("Login: %s", r.GetData().GetName())
	// 	log.Printf("Login: %d", r.GetCode())
	// 	time.Sleep(time.Second * 2)
	// }
	wg.Wait()
}
