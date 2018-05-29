package auth

type AuthServiceType uint32

type AuthServiceParam map[string]string
type AuthServiceData map[string]string

const (
	AST_AUTH         = AuthServiceType(1)
	AST_VERIFICATION = AuthServiceType(2)
	AST_REGISTRY     = AuthServiceType(3)
)

type ClientData struct {
	ClientID     string
	ClientSecret string
}

type AuthToken struct {
	Token          string
	ExpirationTime int64
}

type AuthServiceRequest struct {
	ServiceType AuthServiceType
	Req         AuthRequest
}

func (self *AuthServiceRequest) Init() {
	self.Req.Init()
}

type AuthServiceResponse struct {
	ServiceType AuthServiceType
	Rsp         AuthResponse
}

func (self *AuthServiceResponse) Init() {
	self.Rsp.Init()
}

type AuthRequest struct {
	ReqParam AuthServiceParam
}

func (self *AuthRequest) Init() {
	self.ReqParam = make(AuthServiceParam)
}

type AuthResponse struct {
	Result  bool
	RspData AuthServiceData
}

func (self *AuthResponse) Init() {
	self.RspData = make(AuthServiceData)
}

const (
	AuthTokenValidTime = 1800
)

type ClientAuthTokenMap map[ClientData]AuthToken

type AuthHandlerFunc func(req AuthRequest) (AuthServiceResponse, error)

type AuthHandlerFuncMap map[AuthServiceType]AuthHandlerFunc
