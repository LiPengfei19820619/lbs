package webapi

import (
	"encoding/json"
	"net/http"
	"zte/ims/lbs/securityctrl/bll"
	"zte/ims/lbs/securityctrl/model"
)

// Start 启动WebAPI
func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/selfconfig", handleSelfconfig)

	http.ListenAndServe(":9090", mux)
}

func handleSelfconfig(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		//io.WriteString(w, "selfconfig")
		handleGetSelfConfig(w, req)
	} else {
		http.Error(w, "不支持更新本网元属性配置", 405)
	}

}

func handleGetSelfConfig(w http.ResponseWriter, req *http.Request) {
	var selfConfig model.SelfConfig

	selfConfig.IPAddr = bll.GetSelfIPAddr()
	selfConfig.ListenPort = bll.GetSelfListenPort()
	selfConfig.HostName = "sec.zte.com.cn"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(selfConfig)
}
