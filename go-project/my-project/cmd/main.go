package main

import (
	"log"
	"my-project/internal/adapters/app/api"
	"my-project/internal/adapters/core/arithmetic"
	gRPC "my-project/internal/adapters/framework/left/grpc"
	"my-project/internal/adapters/framework/right/db"
	"my-project/internal/ports"
	"os"
)

func main() {
	var err error
	//ports
	var dbaseAdapter ports.DbPort
	var arithAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	arithAdapter = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, arithAdapter)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
