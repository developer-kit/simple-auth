package zookeeper

type ZKClientData struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type ZKClientDataMap map[string]*ZKClientData
