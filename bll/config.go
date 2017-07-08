package bll

import (
	"fmt"
	"zte/ims/lbs/dal"
	"zte/ims/lbs/model"
)

var (
	selfConfig model.SelfConfig
	adjNodes   []model.AdjacentNode
)

func init() {
	selfConfig = dal.LoadSelfConfig()
	adjNodes = dal.LoadAdjacentNodes()

	fmt.Println("selfconfig:")
	fmt.Println("IPAddr:", selfConfig.IPAddr, ",Port:", selfConfig.ListenPort)

	fmt.Println("adjnodes:")
	for _, node := range adjNodes {
		fmt.Println("NodeID:", node.NodeID, ",NodeType:", node.NodeType, ",IPAddr:", node.IPAddr, ",Port:", node.Port)
	}
}

// GetSelfIPAddr 获取自身的IP地址配置
func GetSelfIPAddr() string {
	return selfConfig.IPAddr
}

// GetSelfListenPort 获取自身监听端口配置
func GetSelfListenPort() int {
	return selfConfig.ListenPort
}

// GetAllAdjacentNodes 获取所有的邻接节点配置
func GetAllAdjacentNodes() []model.AdjacentNode {
	return adjNodes
}

// GetAdjacentNodeByNodeID 根据节点ID获取邻接节点配置
func GetAdjacentNodeByNodeID(nodeID int) *model.AdjacentNode {
	for _, node := range adjNodes {
		if node.NodeID == nodeID {
			return &node
		}
	}

	return nil
}
