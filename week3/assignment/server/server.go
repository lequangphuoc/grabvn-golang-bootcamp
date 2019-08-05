package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grabvn-golang-bootcamp/week3/assignment/passengerfeedback"
)

const (
	port = ":50052"
)

type passengerFeedbackServer struct {
	savedPassengerFeedbacks []*pb.PassengerFeedback
}

func (s *passengerFeedbackServer) AddFeedback(ctx context.Context, in *pb.AddFeedbackRequest) (*pb.AddFeedbackResponse, error) {
	s.savedPassengerFeedbacks = append(s.savedPassengerFeedbacks, in.Feedback)

	return &pb.AddFeedbackResponse {
		ResponseMessage: "Added",
	}, nil
}

func (s *passengerFeedbackServer) GetFeedbacksByPassengerId(ctx context.Context, in *pb.FeedbacksByPassengerIdRequest) (*pb.FeedbacksByPassengerIdResponse, error) {
	result := []*pb.PassengerFeedback{}
	for _, feeback := range s.savedPassengerFeedbacks {
		if feeback.PassengerID == in.PassengerID {
			result = append(result, feeback)
		}
	}

	return &pb.FeedbacksByPassengerIdResponse {
		PassengerFeedbacks: result,
	}, nil
}

func (s *passengerFeedbackServer) GetFeedbacksByBookingCode(ctx context.Context, in *pb.FeedbacksByBookingCodeRequest) (*pb.FeedbacksByBookingCodeResponse, error) {
	result := []*pb.PassengerFeedback{}
	for _, feeback := range s.savedPassengerFeedbacks {
		if feeback.BookingCode == in.BookingCode {
			result = append(result, feeback)
		}
	}

	return &pb.FeedbacksByBookingCodeResponse {
		PassengerFeedbacks: result,
	}, nil
}

func (s *passengerFeedbackServer) DeleteByPassengerId(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	for i, feeback := range s.savedPassengerFeedbacks {
		if feeback.PassengerID == in.PassengerID {
			s.savedPassengerFeedbacks = append(s.savedPassengerFeedbacks[:i], s.savedPassengerFeedbacks[i+1:]...)
		}
	}

	return &pb.DeleteResponse {
		Deleted: true,
	}, nil
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


