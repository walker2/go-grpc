package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"grpc-go-course/calculator/calcpb"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, in *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	log.Printf("Sum function was invoked %v", in)
	var sum int32

	for _, number := range in.GetNumbers().GetNumbers() {
		sum += number
	}

	result := "Result is " + strconv.Itoa(int(sum))
	response := calcpb.SumResponse{
		Result: result,
	}

	return &response, nil
}

func (*server) PrimeNumberDecomposition(req *calcpb.PrimeNumberDecompositionRequest, stream calcpb.SumService_PrimeNumberDecompositionServer) error {
	log.Printf("PrimeNumberDecomposition function was invoked %v", req)
	N := req.GetNumber()
	k := int32(2)

	for N > 1 {
		if N%k == 0 {
			res := &calcpb.PrimeNumberDecompositionResponse{
				Number: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			N = N / k
		} else {
			k++
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calcpb.SumService_ComputeAverageServer) error {
	log.Printf("ComputeAverage function was invoked")
	sum := int32(0)
	n := int32(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			result := float32(sum) / float32(n)
			return stream.SendAndClose(&calcpb.ComputeAverageResponse{
				Number: result,
			})
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
			return err
		}

		sum += req.GetNumber()
		n++
	}
}

func (*server) FindMaximum(stream calcpb.SumService_FindMaximumServer) error {
	log.Printf("FindMaximum function was invoked")
	max := int32(0)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		curr := req.GetNumber()

		if curr > max {
			result := curr
			max = curr
			if err := stream.Send(&calcpb.FindMaximumResponse{Number: result}); err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
				return err
			}
		}

	}
}

func (*server) SquareRoot(ctx context.Context, req *calcpb.SquareRootRequest) (*calcpb.SquareRootResponse, error) {
	log.Printf("SquareRoot function was invoked")
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %v", number),
		)
	}

	return &calcpb.SquareRootResponse{
		NumberRoot: float32(math.Sqrt(float64(number))),
	}, nil
}

func main() {
	fmt.Println("gRPC server is running...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calcpb.RegisterSumServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
