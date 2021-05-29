package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	myFilePath "path/filepath"

	"github.com/LibenHailu/grpc_file_stream/file_stream/filepb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am file stream client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not found connect %v ", err)
	}

	defer cc.Close()

	c := filepb.NewFileServiceClient(cc)

	UploadFile(c, "a", "C:/Users/Liben/Desktop/Liben.jpg")

}

func UploadFile(c filepb.FileServiceClient, fileID string, filepath string) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("couldn't open file: %v", err)
	}

	defer file.Close()

	stream, err := c.UploadFile(context.Background())

	if err != nil {
		log.Fatalf("couldn't upload file %v", err)
	}

	req := &filepb.UploadFileRequest{
		Data: &filepb.UploadFileRequest_Info{
			Info: &filepb.FileInfo{
				FileId:   fileID,
				FileType: myFilePath.Ext(filepath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatalf("couldn't send file %v", err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {

		n, err := reader.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("couldn't read chunk to buffer %v", err)
		}

		req := &filepb.UploadFileRequest{
			Data: &filepb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("couldn't send chunk to server %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("couldn't recive response %v", err)
	}

	log.Printf("file uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}

// func doUnary(c greetpb.GreetServiceClient) {

// 	req := &greetpb.GreetRequest{
// 		Greeting: &greetpb.Greeting{
// 			FirstName: "Liben",
// 			LastName:  "Hailu",
// 		},
// 	}
// 	res, err := c.Greet(context.Background(), req)

// 	if err != nil {
// 		log.Fatalf("error while calling Greet RPC: %v", err)
// 	}

// 	log.Printf("Response form Great: %v", res.Result)
// }
