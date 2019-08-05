package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "grabvn-golang-bootcamp/week3/assignment/passengerfeedback"
)

const (
	address     = "localhost:50052"
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

	// Call Get feedbacks
	req2 := pb.FeedbacksByPassengerIdRequest{
		PassengerID: 1,
	}
	rep2, err := c.GetFeedbacksByPassengerId(ctx, &req2)

	for _, feedback := range(rep2.PassengerFeedbacks) {
		log.Printf("%s", feedback.Feedback)
	}
}