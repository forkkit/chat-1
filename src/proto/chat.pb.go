// Code generated by protoc-gen-go.
// source: chat.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	Chat
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type Chat struct {
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto1.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}

type Chat_Nil struct {
}

func (m *Chat_Nil) Reset()         { *m = Chat_Nil{} }
func (m *Chat_Nil) String() string { return proto1.CompactTextString(m) }
func (*Chat_Nil) ProtoMessage()    {}

type Chat_Packet struct {
	Uid     int32  `protobuf:"varint,1,opt" json:"Uid,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,proto3" json:"Content,omitempty"`
}

func (m *Chat_Packet) Reset()         { *m = Chat_Packet{} }
func (m *Chat_Packet) String() string { return proto1.CompactTextString(m) }
func (*Chat_Packet) ProtoMessage()    {}

type Chat_Id struct {
	Id   int32  `protobuf:"varint,1,opt" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt" json:"Name,omitempty"`
}

func (m *Chat_Id) Reset()         { *m = Chat_Id{} }
func (m *Chat_Id) String() string { return proto1.CompactTextString(m) }
func (*Chat_Id) ProtoMessage()    {}

type Chat_MucReq struct {
	MucId  int32 `protobuf:"varint,1,opt" json:"MucId,omitempty"`
	UserId int32 `protobuf:"varint,2,opt" json:"UserId,omitempty"`
}

func (m *Chat_MucReq) Reset()         { *m = Chat_MucReq{} }
func (m *Chat_MucReq) String() string { return proto1.CompactTextString(m) }
func (*Chat_MucReq) ProtoMessage()    {}

func init() {
}

// Client API for ChatService service

type ChatServiceClient interface {
	Packet(ctx context.Context, opts ...grpc.CallOption) (ChatService_PacketClient, error)
	// Control API
	Reg(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error)
	UpdateInfo(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error)
	RegMuc(ctx context.Context, in *Chat_MucReq, opts ...grpc.CallOption) (*Chat_Nil, error)
	JoinMuc(ctx context.Context, in *Chat_MucReq, opts ...grpc.CallOption) (*Chat_Nil, error)
	LeaveMuc(ctx context.Context, in *Chat_MucReq, opts ...grpc.CallOption) (*Chat_Nil, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Packet(ctx context.Context, opts ...grpc.CallOption) (ChatService_PacketClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatService_serviceDesc.Streams[0], c.cc, "/proto.ChatService/Packet", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServicePacketClient{stream}
	return x, nil
}

type ChatService_PacketClient interface {
	Send(*Chat_Packet) error
	Recv() (*Chat_Packet, error)
	grpc.ClientStream
}

type chatServicePacketClient struct {
	grpc.ClientStream
}

func (x *chatServicePacketClient) Send(m *Chat_Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServicePacketClient) Recv() (*Chat_Packet, error) {
	m := new(Chat_Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) Reg(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/Reg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateInfo(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/UpdateInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RegMuc(ctx context.Context, in *Chat_MucReq, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/RegMuc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinMuc(ctx context.Context, in *Chat_MucReq, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/JoinMuc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) LeaveMuc(ctx context.Context, in *Chat_MucReq, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/LeaveMuc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatService service

type ChatServiceServer interface {
	Packet(ChatService_PacketServer) error
	// Control API
	Reg(context.Context, *Chat_Id) (*Chat_Nil, error)
	UpdateInfo(context.Context, *Chat_Id) (*Chat_Nil, error)
	RegMuc(context.Context, *Chat_MucReq) (*Chat_Nil, error)
	JoinMuc(context.Context, *Chat_MucReq) (*Chat_Nil, error)
	LeaveMuc(context.Context, *Chat_MucReq) (*Chat_Nil, error)
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Packet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Packet(&chatServicePacketServer{stream})
}

type ChatService_PacketServer interface {
	Send(*Chat_Packet) error
	Recv() (*Chat_Packet, error)
	grpc.ServerStream
}

type chatServicePacketServer struct {
	grpc.ServerStream
}

func (x *chatServicePacketServer) Send(m *Chat_Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServicePacketServer) Recv() (*Chat_Packet, error) {
	m := new(Chat_Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_Reg_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Chat_Id)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ChatServiceServer).Reg(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ChatService_UpdateInfo_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Chat_Id)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ChatServiceServer).UpdateInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ChatService_RegMuc_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Chat_MucReq)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ChatServiceServer).RegMuc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ChatService_JoinMuc_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Chat_MucReq)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ChatServiceServer).JoinMuc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ChatService_LeaveMuc_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Chat_MucReq)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ChatServiceServer).LeaveMuc(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reg",
			Handler:    _ChatService_Reg_Handler,
		},
		{
			MethodName: "UpdateInfo",
			Handler:    _ChatService_UpdateInfo_Handler,
		},
		{
			MethodName: "RegMuc",
			Handler:    _ChatService_RegMuc_Handler,
		},
		{
			MethodName: "JoinMuc",
			Handler:    _ChatService_JoinMuc_Handler,
		},
		{
			MethodName: "LeaveMuc",
			Handler:    _ChatService_LeaveMuc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Packet",
			Handler:       _ChatService_Packet_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
