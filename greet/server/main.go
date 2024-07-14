package main
import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "grpctutorial/greet/proto"
	"context"
)

type greetServer struct{
	pb.UnimplementedGreetServiceServer
}
func ( gs greetServer ) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error){
	greetMessage := "GoodMorning "+ in.FirstName
	greetResponse:= &pb.GreetResponse{
		Response : greetMessage,
	}
	return greetResponse,nil
}



func main() {

	// we created a listener
	lis, err := net.Listen("tcp",":50051")
	//we handle error in case there is problem creating the listener
	if err != nil{
		log.Fatalf("unable to create listener, %v")
	}
	//we are creating a grpc server
	s := grpc.NewServer()
	//we are going to register our server with the GreetService
	pb.RegisterGreetServiceServer(s,greetServer{})
	//start the server
	log.Println("Starting gRPC server on port 50051...")
	if err:= s.Serve(lis); err!=nil{
		log.Fatalf("failed to serve: %v", err)
	}



}
