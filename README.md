# Simple Golang gRPC Client/Server example

First you need to create a proto file with version 3 named as "service1.proto" as below:
'''
syntax="proto3";

package service1;

message SensorReadingRep{
	string readings=1;
}

message SensorReadingsReq{
	string sensorno=1;
}

service Sensor{
	rpc SendReadings(SensorReadingsReq) returns (SensorReadingRep) {}
}
'''

After that you need to compile it into Go file using below command:

'''
protoc ./service1.proto --go_out=plugins=grpc:.
'''

*Then you need to create a gRPC server implemtation by doing the following steps:*
1. Create new empty struct which will be dealt with as a Server for your gRPC connectin:
'''
type server struct{}	
'''
2. Listen to a TCP port let it be port 32001 using normal Go "net" packge
3. Create an ew gRPC server usign "grpc.NewServer()" which is defined within "google.golang.org/grpc" Go packge:
'''
s:=grpc.NewServer()
'''
4. Register server to our custom gRPC service:
'''
pb.RegisterSensorServer(s, &server{})
'''
5. Serve clients:
'''
s.Serve(ln)
'''