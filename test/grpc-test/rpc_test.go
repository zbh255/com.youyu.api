package grpc_test

import (
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/test/grpc-test/rpc"
	"com.youyu.api/test/grpc-test/rpc/pb"
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestGrpcInt32Type(t *testing.T) {
	go rpc.Run()
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Error("client cannot dial grpc business_server")
	}
	defer conn.Close()
	client := pb.NewTestApiClient(conn)
	testInt, err := client.TestInt(context.Background(),&pb.TestData{
		Int:  -661,
		Uint: 100,
	})
	if err != nil {
		st,_ := status.FromError(err)
		t.Log(st.Message)
		t.Log(st.Code)
	} else {
		t.Log(testInt.Int)
		t.Log(testInt.Uint)
	}
}
