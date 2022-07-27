package rpc

import (
	"context"
	"log"
	"my-project/internal/adapters/app/api"
	"my-project/internal/adapters/core/arithmetic"
	"my-project/internal/adapters/framework/left/grpc/pb"
	"my-project/internal/adapters/framework/right/db"
	"my-project/internal/ports"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

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

	gRPCAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnecttion(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}
