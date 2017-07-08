package dal

import (
	"database/sql"
	"fmt"
	"zte/ims/lbs/model"
)

var (
	dbconn *sql.DB
)

func init() {
	var err error
	dbconn, err = sql.Open("firebirdsql", "sysdba:root123@127.0.0.1:3050/D:\\tmp\\lbsdata\\LBS_CONFIG.FDB")
	if err != nil {
		fmt.Println("connect to db failed, err: ", err.Error())
	}
}

// LoadSelfConfig 从数据库中加载本网元自身配置信息
func LoadSelfConfig() model.SelfConfig {
	var selfConfig model.SelfConfig

	rows, err := dbconn.Query("SELECT IPADDR, PORT FROM R_SELFCONFIG;")
	if err != nil {
		fmt.Println("query r_selfconfig table failed, err: ", err.Error())
		return selfConfig
	}

	for rows.Next() {
		err = rows.Scan(&selfConfig.IPAddr, &selfConfig.ListenPort)
		if err == nil {
			break
		}
	}

	/* selfConfig.IPAddr = "10.43.31.148"
	selfConfig.ListenPort = 3001 */

	return selfConfig
}

// LoadAdjacentNodes 加载邻接节点配置信息
func LoadAdjacentNodes() []model.AdjacentNode {
	adjNodes := make([]model.AdjacentNode, 0)

	rows, err := dbconn.Query("SELECT NODEID, NODETYPE, IPADDR, PORT FROM R_ADJNODE;")
	if err != nil {
		fmt.Println("query r_adjnode table failed, err: ", err.Error())
		return adjNodes
	}

	for rows.Next() {
		var node model.AdjacentNode
		err = rows.Scan(&node.NodeID, &node.NodeType, &node.IPAddr, &node.Port)
		if err != nil {
			fmt.Println("scan r_adjnode row failed, err: ", err.Error())
			break
		}

		adjNodes = append(adjNodes, node)
	}

	/* node := model.AdjacentNode{NodeID: 1, NodeType: 0, IPAddr: "10.43.154.140", Port: 3002}
	adjNodes = append(adjNodes, node)

	node = model.AdjacentNode{NodeID: 2, NodeType: 0, IPAddr: "10.43.154.140", Port: 3003}
	adjNodes = append(adjNodes, node) */

	return adjNodes
}
