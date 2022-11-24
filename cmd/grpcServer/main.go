package main

import (
	"database/sql"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/jonasjesusamerico/full-cycle-grpc-golang/database"
	"github.com/jonasjesusamerico/full-cycle-grpc-golang/internal/pb"
	"github.com/jonasjesusamerico/full-cycle-grpc-golang/internal/pb/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	// Cria o servidor GRPC
	grpcServer := grpc.NewServer()

	// É necessário registrar o servidor no serviço
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// Necessário registra no reflection para funiconamento correto do cliente
	reflection.Register(grpcServer)

	// Abre porta TPC para ter comunicação com o GRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
