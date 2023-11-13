package api

import (
	"fmt"
	"github.com/451008604/nets/iface"
	"google.golang.org/protobuf/proto"
	"netsFrame/logic"
	pb "netsFrame/proto/bin"
)

func init() {
	// 注册路由
	RegisterRouter()
}

func MsgFilter(c iface.IRequest, req proto.Message) bool {
	// 未登录时不处理任何请求
	if c.GetMsgID() != int32(pb.MSgID_PlayerLogin_Req) && logic.GetPlayer(c.GetConnection()).Data == nil {
		return false
	}

	return true
}

func MsgErrCapture(request iface.IRequest, r any) {
	request.GetConnection().SendMsg(int32(pb.MSgID_ServerErr_Notify), &pb.ServerErrNotify{
		ErrCode: 1,
		ErrMsg:  fmt.Sprintf("%v", r),
	})
}
