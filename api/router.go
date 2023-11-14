package api

import (
	"github.com/451008604/nets/iface"
	"github.com/451008604/nets/network"
	"google.golang.org/protobuf/proto"
	pb "netsFrame/proto/bin"
)

func RegisterRouter() {
	msg := network.GetInstanceMsgHandler()
	msg.SetFilter(MsgFilter)
	msg.SetErrCapture(MsgErrCapture)
	RegisterRouter2(msg, pb.MSgID_Heartbeat_Req, func() proto.Message { return new(pb.HeartbeatRequest) }, HeartBeatHandler)
	RegisterRouter2(msg, pb.MSgID_Heartbeat_Req, func() proto.Message { return new(pb.HeartbeatRequest) }, HeartBeatHandler)
	RegisterRouter2(msg, pb.MSgID_PlayerLogin_Req, func() proto.Message { return new(pb.PlayerLoginRequest) }, LoginHandler)
	RegisterRouter2(msg, pb.MSgID_Broadcast_Req, func() proto.Message { return new(pb.BroadcastRequest) }, BroadcastHandler)
}

func RegisterRouterClient() {
	msg := network.GetInstanceMsgHandler()
	RegisterRouterClient2(msg, pb.MSgID_Heartbeat_Res, func() proto.Message { return new(pb.HeartbeatResponse) })
	RegisterRouterClient2(msg, pb.MSgID_PlayerLogin_Res, func() proto.Message { return new(pb.PlayerLoginResponse) })
	RegisterRouterClient2(msg, pb.MSgID_Broadcast_Res, func() proto.Message { return new(pb.BroadcastResponse) })
	RegisterRouterClient2(msg, pb.MSgID_ServerErr_Notify, func() proto.Message { return new(pb.ServerErrNotify) })

}

func RegisterRouter2(a iface.IMsgHandler, b pb.MSgID, c func() proto.Message, d func(c iface.IConnection, _ proto.Message)) {
	a.AddRouter(int32(b), c, d)
}
func RegisterRouterClient2(a iface.IMsgHandler, b pb.MSgID, c func() proto.Message) {
	a.AddRouter(int32(b), c, nil)
}
