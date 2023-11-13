package api

import (
	"github.com/451008604/nets/network"
	"google.golang.org/protobuf/proto"
	pb "netsFrame/proto/bin"
)

func RegisterRouter() {
	msgHandler := network.GetInstanceMsgHandler()
	msgHandler.SetFilter(MsgFilter)
	msgHandler.SetErrCapture(MsgErrCapture)
	msgHandler.AddRouter(int32(pb.MSgID_Heartbeat_Req), func() proto.Message { return new(pb.HeartbeatRequest) }, HeartBeatHandler)
	msgHandler.AddRouter(int32(pb.MSgID_PlayerLogin_Req), func() proto.Message { return new(pb.PlayerLoginRequest) }, LoginHandler)
	msgHandler.AddRouter(int32(pb.MSgID_Broadcast_Req), func() proto.Message { return new(pb.BroadcastRequest) }, BroadcastHandler)
}

func RegisterRouterClient() {
	msgHandler := network.GetInstanceMsgHandler()
	msgHandler.AddRouter(int32(pb.MSgID_Heartbeat_Res), func() proto.Message { return new(pb.HeartbeatResponse) }, nil)
	msgHandler.AddRouter(int32(pb.MSgID_PlayerLogin_Res), func() proto.Message { return new(pb.PlayerLoginResponse) }, nil)
	msgHandler.AddRouter(int32(pb.MSgID_Broadcast_Res), func() proto.Message { return new(pb.BroadcastResponse) }, nil)
	msgHandler.AddRouter(int32(pb.MSgID_ServerErr_Notify), func() proto.Message { return new(pb.ServerErrNotify) }, nil)
}
