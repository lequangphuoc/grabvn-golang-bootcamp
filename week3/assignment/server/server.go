package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grabvn-golang-bootcamp/week3/assignment/passengerfeedback"
)

const (
	port = ":50051"
)

type passengerFeedbackServer struct {
	savedPassengerFeedback []*pb.PassengerFeedback
}

func (s *passengerFeedbackServer) AddFeedback(ctx context.Context, in *pb.AddFeedbackRequest) (*pb.AddFeedbackResponse, error) {
	return nil, nil
}

func (s *passengerFeedbackServer) GetFeedbacksByPassengerId(ctx context.Context, in *pb.FeedbacksByPassengerIdRequest) (*pb.FeedbacksByPassengerIdResponse, error) {
	return nil, nil
}

func (s *passengerFeedbackServer) GetFeedbacksByBookingCode(ctx context.Context, in *pb.FeedbacksByBookingCodeRequest) (*pb.FeedbacksByBookingCodeResponse, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer  := grpc.NewServer()
	pb.RegisterPassengerFeedbackServiceServer(grpcServer , &passengerFeedbackServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


