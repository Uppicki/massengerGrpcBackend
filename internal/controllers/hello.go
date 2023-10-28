package controllers

import (
	"context"
	"fmt"
)

type HelloService struct {
	UnimplementedTesterServer
}

func (h *HelloService) Test(ctx context.Context, request *TestRequest) (*TestResponse, error) {
	fmt.Println(request.Request)

	return &TestResponse{
		Response: string("Asd"),
	}, nil
}
