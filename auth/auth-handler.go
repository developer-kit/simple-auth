package auth

func handleAuthRequest(req AuthRequest) (AuthServiceResponse, error) {
	rsp := AuthServiceResponse{}
	rsp.Init()
	rsp.Rsp.Result = true
	rsp.Rsp.RspData["Result"] = "true"
	if GetInstance().CheckClientSecret(req.ReqParam["ClientID"], req.ReqParam["ClientSecret"]) {
		rsp.Rsp.RspData["Token"] = GetInstance().GetAuthTokenManager().GetAuthToken(req.ReqParam["ClientID"], req.ReqParam["ClientSecret"])
	} else {
		rsp.Rsp.RspData["Token"] = ""
	}
	return rsp, nil
}

func handleVerificationRequest(req AuthRequest) (AuthServiceResponse, error) {
	rsp := AuthServiceResponse{}
	rsp.Init()
	rsp.Rsp.Result = true
	resStr := "true"
	res, clientID := GetInstance().GetAuthTokenManager().VerificationAuthToken(req.ReqParam["Token"])
	if !res {
		resStr = "false"
	}
	rsp.Rsp.RspData["Result"] = resStr
	rsp.Rsp.RspData["ClientID"] = clientID
	return rsp, nil
}

func handleRegistryRequest(req AuthRequest) (AuthServiceResponse, error) {
	rsp := AuthServiceResponse{}
	rsp.Init()
	rsp.Rsp.Result = true
	err := GetInstance().RegistryClient(req.ReqParam["ClientID"], req.ReqParam["ClientSecret"])
	resStr := "true"
	if nil != err {
		resStr = "false"
	}
	rsp.Rsp.RspData["Result"] = resStr
	return rsp, nil
}
