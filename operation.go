package mandela

import (
	"encoding/json"
	"fmt"
	msg "github.com/prestonTao/mandela/message"
	"github.com/prestonTao/mandela/net"
	"github.com/prestonTao/mandela/nodeStore"
)

//保存一个键值对
func SaveData(key, value string) {
	clientConn, _ := engine.GetController().GetSession(nodeManager.SuperName)
	data := []byte(key + "!" + value)
	clientConn.Send(msg.SaveKeyValueReqNum, &data)
}

//给所有客户端发送消息
func SendMsgForAll(message string) {
	messageSend := msg.Message{
		Content: []byte(message),
	}
	for idOne, nodeOne := range nodeManager.GetAllNodes() {
		if clientConn, ok := engine.GetController().GetSession(string(nodeOne.IdInfo.Build())); ok {
			messageSend.TargetId = idOne
			data, _ := json.Marshal(messageSend)
			clientConn.Send(msg.SendMessage, &data)
		}
	}
}

//给某个人发送消息
func SendMsgForOne(target, message string) {
	//发送给自己的
	if nodeStore.ParseId(nodeManager.GetRootIdInfoString()) == target {
		fmt.Println(message)
		return
	}
	messageSend := msg.Message{
		TargetId: target,
		Content:  []byte(message),
	}
	sendBytes, _ := json.Marshal(&messageSend)

	var session net.Session
	var ok bool
	//本机是超级节点
	if Init_IsSuperPeer {
		//是发给自己的弱节点
		if targetNode, ok := nodeManager.GetProxyNode(target); ok {
			if session, ok := engine.GetController().GetSession(string(targetNode.IdInfo.Build())); ok {
				err := session.Send(msg.SendMessage, &sendBytes)
				if err != nil {
					fmt.Println("message发送数据出错：", err.Error())
				}
			} else {
				//这个节点离线了，想办法处理下
			}
			return
		}
		//转发出去
		targetNode := nodeManager.Get(target, true, "")
		if targetNode == nil {
			fmt.Println("本机未连入mandela网络")
			return
		}
		session, ok = engine.GetController().GetSession(string(targetNode.IdInfo.Build()))
	} else {
		//本机是普通节点
		//获得超级节点
		session, ok = engine.GetController().GetSession(nodeManager.SuperName)
	}
	if !ok {
		return
	}
	err := session.Send(msg.SendMessage, &sendBytes)
	if err != nil {
		fmt.Println("message发送数据出错：", err.Error())
	}
}

//注册一个域名帐号
func CreateAccount(account string) {
	// id := GetHashKey(account)
}

/*
	查看本地保存的所有节点id
*/
func See() {
	allNodes := nodeManager.GetAllNodes()
	for key, _ := range allNodes {
		fmt.Println(key)
	}
}

/*
	查看本地保存的节点中，小于本节点id的所有节点
*/
func SeeLeftNode() {
	nodes := nodeManager.GetLeftNode(*nodeManager.Root.IdInfo.GetBigIntId(), nodeManager.MaxRecentCount)
	for _, id := range nodes {
		fmt.Println(id.IdInfo.GetId())
	}
}

/*
	查看本地保存的节点中，大于本节点id的所有节点
*/
func SeeRightNode() {
	nodes := nodeManager.GetRightNode(*nodeManager.Root.IdInfo.GetBigIntId(), nodeManager.MaxRecentCount)
	for _, id := range nodes {
		fmt.Println(id.IdInfo.GetId())
	}
}

/*
	创建一个id
*/
func CreateIdInfo() {

}
