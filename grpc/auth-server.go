package grpc

import (
	"errors"
	"github.com/seaofstars-coder/simple-auth/auth"
	"github.com/seaofstars-coder/simple-auth/proto"
	"golang.org/x/net/context"
)

type AuthGrpcServer struct {
}

func (self *AuthGrpcServer) DoAuth(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	request := auth.AuthServiceRequest{}
	request.Init()
	request.ServiceType = auth.AST_AUTH
	request.Req.ReqParam["ClientID"] = req.ClientId
	request.Req.ReqParam["ClientSecret"] = req.ClientSecret
	authResponse := auth.GetInstance().Handle(request)
	ret := proto.AuthResponse_SUCCESS
	var err error = nil
	if !authResponse.Rsp.Result {
		ret = proto.AuthResponse_FAIL
		err = errors.New("handle DoAuth fail!")
	}
	return &proto.AuthResponse{ret, authResponse.Rsp.RspData["Token"]}, err
}
