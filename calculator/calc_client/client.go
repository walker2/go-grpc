package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calcpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC client is running...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer cc.Close()

	c := calcpb.NewSumServiceClient(cc)

	//doSum(c)
	//doPrimeNumberDecomposition(c)
	//doComputeAverage(c)
	//doFindMax(c)
	doErrorUnary(c)
}
func doErrorUnary(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do Square Root RPC call...")
	number := int32(10)
	request := &calcpb.SquareRootRequest{Number: number}
	resp, err := c.SquareRoot(context.Background(), request)

	if err != nil {
		respErr, ok := status.FromError(err)

		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent negative number!")
			}
			return
		}
		log.Fatalf("Big error in square root %v", err)
	}

	log.Printf("Result of sqrt of %v is %v", number, resp.GetNumberRoot())
}

func doFindMax(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do BiDi streaming RPC...")

	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while calling greet everyone: %v", err)
		return
	}
	waitc := make(chan struct{})

	requests := []*calcpb.FindMaximumRequest{
		&calcpb.FindMaximumRequest{
			Number: 1,
		},
		&calcpb.FindMaximumRequest{
			Number: 5,
		},
		&calcpb.FindMaximumRequest{
			Number: 3,
		},
		&calcpb.FindMaximumRequest{
			Number: 6,
		},
		&calcpb.FindMaximumRequest{
			Number: 2,
		},
		&calcpb.FindMaximumRequest{
			Number: 20,
		},
	}

	go func() {
		// Function to send a buch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetNumber())
		}
		close(waitc)
	}()

	// Block until everything done
	<-waitc
}

func doComputeAverage(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do Client streaming RPC...")
	request := []*calcpb.ComputeAverageRequest{
		&calcpb.ComputeAverageRequest{
			Number: 2,
		},
		&calcpb.ComputeAverageRequest{
			Number: 8,
		},
		&calcpb.ComputeAverageRequest{
			Number: 10,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling compute average: %v", err)
	}

	for _, req := range request {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response compute average: %v", err)
	}
	fmt.Printf("Compute average Response: %v\n", res)
}

func doPrimeNumberDecomposition(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do Server streaming RPC...")
	request := &calcpb.PrimeNumberDecompositionRequest{
		Number: 65492,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling primenumber rpc %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// We've reached end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Printf("Response from primenumber: %v", msg.GetNumber())
	}
}

func doSum(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do Unary RPC call...")

	request := &calcpb.SumRequest{
		Numbers: &calcpb.Numbers{
			Numbers: []int32{5, 7, 3, 15},
		},
	}
	resp, err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling greet rpc %v", err)
	}

	log.Printf("Response from calc %v", resp.Result)
}
