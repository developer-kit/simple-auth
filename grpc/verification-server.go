package grpc

import (
	"errors"
	"github.com/seaofstars-coder/simple-auth/auth"
	"github.com/seaofstars-coder/simple-auth/proto"
	"golang.org/x/net/context"
)

type VerificationGrpcServer struct {
}

func (self *VerificationGrpcServer) DoVerification(ctx context.Context, req *proto.VerificationRequest) (*proto.VerificationResponse, error) {
	request := auth.AuthServiceRequest{}
	request.Init()
	request.ServiceType = auth.AST_VERIFICATION
	request.Req.ReqParam["Token"] = req.Token
	authResponse := auth.GetInstance().Handle(request)
	ret := proto.VerificationResponse_SUCCESS
	clientID := ""
	var err error = nil
	if !authResponse.Rsp.Result {
		ret = proto.VerificationResponse_FAIL
		err = errors.New("handle DoVerification fail!")
	}
	if "false" == authResponse.Rsp.RspData["Result"] {
		ret = proto.VerificationResponse_FAIL
		err = errors.New("verification false")
	}
	clientID = authResponse.Rsp.RspData["ClientID"]
	return &proto.VerificationResponse{ret, clientID}, err
}
