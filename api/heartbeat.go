package api

import (
	"github.com/451008604/nets/iface"
	"google.golang.org/protobuf/proto"
	pb "netsFrame/proto/bin"
	"time"
)

func HeartBeatHandler(c iface.IConnection, _ proto.Message) {
	// 发送给客户端的数据
	c.SendMsg(int32(pb.MSgID_Heartbeat_Res), &pb.HeartbeatResponse{ServerTime: uint32(time.Now().Unix())})
}
