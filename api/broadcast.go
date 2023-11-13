package api

import (
	"github.com/451008604/nets/iface"
	"github.com/451008604/nets/network"
	"google.golang.org/protobuf/proto"
	pb "netsFrame/proto/bin"
)

func BroadcastHandler(c iface.IConnection, req proto.Message) {
	msgReq := req.(*pb.BroadcastRequest)
	msgRes := &pb.BroadcastResponse{
		Result: 0,
		Str:    msgReq.GetStr(),
	}

	network.GetInsBroadcastManager().GetGlobalBroadcast().BroadcastAllTargets(c.GetConnID(), int32(pb.MSgID_Broadcast_Res), msgRes)

	msgRes.Str = "广播成功 " + msgRes.GetStr()
	// c.SendMsg(pb.MSgID_Broadcast_Res, msgRes)
}
