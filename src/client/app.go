package main

import(
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "service1"
)

func main() {
	conn,err:= grpc.Dial("localhost:32001", grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("%s\n", err)
	}
	defer conn.Close()
	c:= pb.NewSensorClient(conn)

	r,err:= c.SendReadings(context.Background(), &pb.SensorReadingsReq{Sensorno:"2"})
	if err!=nil{
		log.Fatalf("%s\n",err)
	}

	log.Printf("Sensor Readings: %s\n", r.Readings)
}