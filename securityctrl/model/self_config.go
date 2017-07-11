package model

// SelfConfig 本网元自身配置信息
type SelfConfig struct {
	IPAddr     string `json:"ipaddr"`
	ListenPort int    `json:"listenport"`
	HostName   string `json:"hostname"`
}
