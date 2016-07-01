package main

import(
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "service1"

	"log"
	"net"
)

type server struct{}

func (s *server) SendReadings(ctx context.Context, r *pb.SensorReadingsReq) (*pb.SensorReadingRep, error) {
	return &pb.SensorReadingRep{Readings:"222"},nil	
}

func main() {
	ln, err:= net.Listen("tcp", ":32001")
	if err!=nil{
		log.Fatalf("%s\n",err)
	}
	//First Create a new gRPC Server 
	s:=grpc.NewServer()
	//Reigster our new server as gRPC
	pb.RegisterSensorServer(s, &server{})
	//Serve our new clients
	s.Serve(ln)
}