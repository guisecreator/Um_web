package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	service "github.com/guisecreator/um_backend/proto"
)

type server struct {
	service.UnimplementedAddServiceServer
	db *pgx.Conn
}

func main() {
	conn, err := pgx.Connect(context.Background(),
		"postgres://postgres:0000@localhost:5432/um_service1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			fprintf, err := fmt.Fprintf(
				os.Stderr, "Unable to close connection: %v\n", err)
			if err != nil {
				panic(fprintf)
				return
			}
			os.Exit(1)
		}
	}(conn, context.Background())

	//db := db.InitBunDb()

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Println(err)
	}

	serv := grpc.NewServer()

	service.RegisterAddServiceServer(serv, &server{db: conn})
	reflection.Register(serv)
	if e := serv.Serve(lis); e != nil {
		log.Println(e)
	}
}
