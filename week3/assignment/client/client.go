package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grabvn-golang-bootcamp/week3/assignment/passengerfeedback"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPassengerFeedbackServiceClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call Create
	req1 := pb.AddFeedbackRequest{
		Feedback: &pb.PassengerFeedback{
			BookingCode: "BC001",
			PassengerID: 1,
			Feedback: "Very good",
		},
	}
	rep1, err := c.AddFeedback(ctx, &req1)

	log.Printf("%s", rep1.ResponseMessage)
}