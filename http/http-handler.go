package http

import (
	"encoding/json"
	"github.com/seaofstars-coder/simple-auth/auth"
	"net/http"
)

func handleAuthRequest(w http.ResponseWriter, r *http.Request) {
	request := auth.AuthServiceRequest{}
	request.Init()
	request.ServiceType = auth.AST_AUTH
	r.ParseForm()
	if http.MethodGet == r.Method {
		request.Req.ReqParam["ClientID"] = r.Form.Get("ClientID")
		request.Req.ReqParam["ClientSecret"] = r.Form.Get("ClientSecret")
	} else if http.MethodPost == r.Method {
		request.Req.ReqParam["ClientID"] = r.PostForm.Get("ClientID")
		request.Req.ReqParam["ClientSecret"] = r.PostForm.Get("ClientSecret")
	}
	authResponse := auth.GetInstance().Handle(request)
	jsonData, err := json.Marshal(authResponse.Rsp)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for k, v := range authResponse.Rsp.RspData {
		w.Header().Set(k, v)
	}
	if authResponse.Rsp.Result {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func handleVerificationRequest(w http.ResponseWriter, r *http.Request) {
	request := auth.AuthServiceRequest{}
	request.Init()
	request.ServiceType = auth.AST_VERIFICATION
	r.ParseForm()
	if http.MethodGet == r.Method {
		request.Req.ReqParam["ClientID"] = r.Form.Get("ClientID")
		request.Req.ReqParam["ClientSecret"] = r.Form.Get("ClientSecret")
		request.Req.ReqParam["Token"] = r.Form.Get("Token")
	} else if http.MethodPost == r.Method {
		request.Req.ReqParam["ClientID"] = r.PostForm.Get("ClientID")
		request.Req.ReqParam["ClientSecret"] = r.PostForm.Get("ClientSecret")
		request.Req.ReqParam["Token"] = r.PostForm.Get("Token")
	}
	authResponse := auth.GetInstance().Handle(request)
	jsonData, err := json.Marshal(authResponse.Rsp)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for k, v := range authResponse.Rsp.RspData {
		w.Header().Set(k, v)
	}
	if authResponse.Rsp.Result {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func handleRegistryRequest(w http.ResponseWriter, r *http.Request) {
	clientID := ""
	clientSecret := ""
	r.ParseForm()
	if http.MethodPost == r.Method {
		clientID = r.PostForm.Get("ClientID")
		clientSecret = r.PostForm.Get("ClientSecret")
	} else if http.MethodGet == r.Method {
		clientID = r.Form.Get("ClientID")
		clientSecret = r.Form.Get("ClientSecret")
	}
	if "" == clientID || "" == clientSecret {
		GetInstance().outputHTML(w, r, "registry.html")
	} else {
		request := auth.AuthServiceRequest{}
		request.Init()
		request.ServiceType = auth.AST_REGISTRY
		request.Req.ReqParam["ClientID"] = clientID
		request.Req.ReqParam["ClientSecret"] = clientSecret
		authResponse := auth.GetInstance().Handle(request)
		jsonData, err := json.Marshal(authResponse.Rsp)
		if nil != err {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for k, v := range authResponse.Rsp.RspData {
			w.Header().Set(k, v)
		}
		if authResponse.Rsp.Result {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(jsonData)
	}
}
