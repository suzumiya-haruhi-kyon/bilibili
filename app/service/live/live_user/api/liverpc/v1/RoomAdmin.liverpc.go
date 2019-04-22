// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/RoomAdmin.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ===================
// RoomAdmin Interface
// ===================

type RoomAdmin interface {
	// * 判断是否房管,防止和room循环调用,传入所有参数
	//
	IsRoomAdmin(context.Context, *RoomAdminIsRoomAdminReq) (*RoomAdminIsRoomAdminResp, error)
}

// =========================
// RoomAdmin Live Rpc Client
// =========================

type roomAdminRpcClient struct {
	client *liverpc.Client
}

// NewRoomAdminRpcClient creates a Rpc client that implements the RoomAdmin interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewRoomAdminRpcClient(client *liverpc.Client) RoomAdmin {
	return &roomAdminRpcClient{
		client: client,
	}
}

func (c *roomAdminRpcClient) IsRoomAdmin(ctx context.Context, in *RoomAdminIsRoomAdminReq) (*RoomAdminIsRoomAdminResp, error) {
	out := new(RoomAdminIsRoomAdminResp)
	err := doRpcRequest(ctx, c.client, 1, "RoomAdmin.is_room_admin", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}