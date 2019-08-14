// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_msg.proto

package pb_msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//--0
type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

//--1
type Pong struct {
	ServerTime           int64    `protobuf:"varint,1,opt,name=serverTime,proto3" json:"serverTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetServerTime() int64 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

//--2
type ErrMsg_S2C struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrMsg_S2C) Reset()         { *m = ErrMsg_S2C{} }
func (m *ErrMsg_S2C) String() string { return proto.CompactTextString(m) }
func (*ErrMsg_S2C) ProtoMessage()    {}
func (*ErrMsg_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{2}
}

func (m *ErrMsg_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrMsg_S2C.Unmarshal(m, b)
}
func (m *ErrMsg_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrMsg_S2C.Marshal(b, m, deterministic)
}
func (m *ErrMsg_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrMsg_S2C.Merge(m, src)
}
func (m *ErrMsg_S2C) XXX_Size() int {
	return xxx_messageInfo_ErrMsg_S2C.Size(m)
}
func (m *ErrMsg_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrMsg_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_ErrMsg_S2C proto.InternalMessageInfo

func (m *ErrMsg_S2C) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ErrMsg_S2C) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ErrMsg_S2C) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

//--3
type LoginInfo_C2S struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	GameId               string   `protobuf:"bytes,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
	ServerUrl            string   `protobuf:"bytes,3,opt,name=serverUrl,proto3" json:"serverUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginInfo_C2S) Reset()         { *m = LoginInfo_C2S{} }
func (m *LoginInfo_C2S) String() string { return proto.CompactTextString(m) }
func (*LoginInfo_C2S) ProtoMessage()    {}
func (*LoginInfo_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{3}
}

func (m *LoginInfo_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginInfo_C2S.Unmarshal(m, b)
}
func (m *LoginInfo_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginInfo_C2S.Marshal(b, m, deterministic)
}
func (m *LoginInfo_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginInfo_C2S.Merge(m, src)
}
func (m *LoginInfo_C2S) XXX_Size() int {
	return xxx_messageInfo_LoginInfo_C2S.Size(m)
}
func (m *LoginInfo_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_LoginInfo_C2S proto.InternalMessageInfo

func (m *LoginInfo_C2S) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LoginInfo_C2S) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *LoginInfo_C2S) GetServerUrl() string {
	if m != nil {
		return m.ServerUrl
	}
	return ""
}

type LoginData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	HeadImg              string   `protobuf:"bytes,3,opt,name=headImg,proto3" json:"headImg,omitempty"`
	Account              float64  `protobuf:"fixed64,4,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginData) Reset()         { *m = LoginData{} }
func (m *LoginData) String() string { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()    {}
func (*LoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{4}
}

func (m *LoginData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginData.Unmarshal(m, b)
}
func (m *LoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginData.Marshal(b, m, deterministic)
}
func (m *LoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginData.Merge(m, src)
}
func (m *LoginData) XXX_Size() int {
	return xxx_messageInfo_LoginData.Size(m)
}
func (m *LoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginData proto.InternalMessageInfo

func (m *LoginData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LoginData) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *LoginData) GetHeadImg() string {
	if m != nil {
		return m.HeadImg
	}
	return ""
}

func (m *LoginData) GetAccount() float64 {
	if m != nil {
		return m.Account
	}
	return 0
}

//--4
type LoginInfo_S2C struct {
	LoginData            *LoginData `protobuf:"bytes,1,opt,name=loginData,proto3" json:"loginData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginInfo_S2C) Reset()         { *m = LoginInfo_S2C{} }
func (m *LoginInfo_S2C) String() string { return proto.CompactTextString(m) }
func (*LoginInfo_S2C) ProtoMessage()    {}
func (*LoginInfo_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{5}
}

func (m *LoginInfo_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginInfo_S2C.Unmarshal(m, b)
}
func (m *LoginInfo_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginInfo_S2C.Marshal(b, m, deterministic)
}
func (m *LoginInfo_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginInfo_S2C.Merge(m, src)
}
func (m *LoginInfo_S2C) XXX_Size() int {
	return xxx_messageInfo_LoginInfo_S2C.Size(m)
}
func (m *LoginInfo_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_LoginInfo_S2C proto.InternalMessageInfo

func (m *LoginInfo_S2C) GetLoginData() *LoginData {
	if m != nil {
		return m.LoginData
	}
	return nil
}

//--5
type JoinRoom_C2S struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoom_C2S) Reset()         { *m = JoinRoom_C2S{} }
func (m *JoinRoom_C2S) String() string { return proto.CompactTextString(m) }
func (*JoinRoom_C2S) ProtoMessage()    {}
func (*JoinRoom_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{6}
}

func (m *JoinRoom_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoom_C2S.Unmarshal(m, b)
}
func (m *JoinRoom_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoom_C2S.Marshal(b, m, deterministic)
}
func (m *JoinRoom_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoom_C2S.Merge(m, src)
}
func (m *JoinRoom_C2S) XXX_Size() int {
	return xxx_messageInfo_JoinRoom_C2S.Size(m)
}
func (m *JoinRoom_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoom_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoom_C2S proto.InternalMessageInfo

func (m *JoinRoom_C2S) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type RoomData struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomData) Reset()         { *m = RoomData{} }
func (m *RoomData) String() string { return proto.CompactTextString(m) }
func (*RoomData) ProtoMessage()    {}
func (*RoomData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{7}
}

func (m *RoomData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomData.Unmarshal(m, b)
}
func (m *RoomData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomData.Marshal(b, m, deterministic)
}
func (m *RoomData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomData.Merge(m, src)
}
func (m *RoomData) XXX_Size() int {
	return xxx_messageInfo_RoomData.Size(m)
}
func (m *RoomData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomData.DiscardUnknown(m)
}

var xxx_messageInfo_RoomData proto.InternalMessageInfo

func (m *RoomData) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

//--6
type JoinRoom_S2C struct {
	RoomData             *RoomData `protobuf:"bytes,1,opt,name=roomData,proto3" json:"roomData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *JoinRoom_S2C) Reset()         { *m = JoinRoom_S2C{} }
func (m *JoinRoom_S2C) String() string { return proto.CompactTextString(m) }
func (*JoinRoom_S2C) ProtoMessage()    {}
func (*JoinRoom_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{8}
}

func (m *JoinRoom_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoom_S2C.Unmarshal(m, b)
}
func (m *JoinRoom_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoom_S2C.Marshal(b, m, deterministic)
}
func (m *JoinRoom_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoom_S2C.Merge(m, src)
}
func (m *JoinRoom_S2C) XXX_Size() int {
	return xxx_messageInfo_JoinRoom_S2C.Size(m)
}
func (m *JoinRoom_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoom_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoom_S2C proto.InternalMessageInfo

func (m *JoinRoom_S2C) GetRoomData() *RoomData {
	if m != nil {
		return m.RoomData
	}
	return nil
}

//--7
type LeaveRoom_C2S struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoom_C2S) Reset()         { *m = LeaveRoom_C2S{} }
func (m *LeaveRoom_C2S) String() string { return proto.CompactTextString(m) }
func (*LeaveRoom_C2S) ProtoMessage()    {}
func (*LeaveRoom_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{9}
}

func (m *LeaveRoom_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoom_C2S.Unmarshal(m, b)
}
func (m *LeaveRoom_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoom_C2S.Marshal(b, m, deterministic)
}
func (m *LeaveRoom_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoom_C2S.Merge(m, src)
}
func (m *LeaveRoom_C2S) XXX_Size() int {
	return xxx_messageInfo_LeaveRoom_C2S.Size(m)
}
func (m *LeaveRoom_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoom_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoom_C2S proto.InternalMessageInfo

//--8
type LeaveRoom_S2C struct {
	LoginData            *LoginData `protobuf:"bytes,1,opt,name=loginData,proto3" json:"loginData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LeaveRoom_S2C) Reset()         { *m = LeaveRoom_S2C{} }
func (m *LeaveRoom_S2C) String() string { return proto.CompactTextString(m) }
func (*LeaveRoom_S2C) ProtoMessage()    {}
func (*LeaveRoom_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{10}
}

func (m *LeaveRoom_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoom_S2C.Unmarshal(m, b)
}
func (m *LeaveRoom_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoom_S2C.Marshal(b, m, deterministic)
}
func (m *LeaveRoom_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoom_S2C.Merge(m, src)
}
func (m *LeaveRoom_S2C) XXX_Size() int {
	return xxx_messageInfo_LeaveRoom_S2C.Size(m)
}
func (m *LeaveRoom_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoom_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoom_S2C proto.InternalMessageInfo

func (m *LeaveRoom_S2C) GetLoginData() *LoginData {
	if m != nil {
		return m.LoginData
	}
	return nil
}

//--9
type PlayerAction_C2S struct {
	Bet                  float64  `protobuf:"fixed64,1,opt,name=bet,proto3" json:"bet,omitempty"`
	Pot                  int32    `protobuf:"varint,2,opt,name=pot,proto3" json:"pot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerAction_C2S) Reset()         { *m = PlayerAction_C2S{} }
func (m *PlayerAction_C2S) String() string { return proto.CompactTextString(m) }
func (*PlayerAction_C2S) ProtoMessage()    {}
func (*PlayerAction_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{11}
}

func (m *PlayerAction_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerAction_C2S.Unmarshal(m, b)
}
func (m *PlayerAction_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerAction_C2S.Marshal(b, m, deterministic)
}
func (m *PlayerAction_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAction_C2S.Merge(m, src)
}
func (m *PlayerAction_C2S) XXX_Size() int {
	return xxx_messageInfo_PlayerAction_C2S.Size(m)
}
func (m *PlayerAction_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAction_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAction_C2S proto.InternalMessageInfo

func (m *PlayerAction_C2S) GetBet() float64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *PlayerAction_C2S) GetPot() int32 {
	if m != nil {
		return m.Pot
	}
	return 0
}

//--10
type PlayerAction_S2C struct {
	RoomData             *RoomData `protobuf:"bytes,1,opt,name=roomData,proto3" json:"roomData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlayerAction_S2C) Reset()         { *m = PlayerAction_S2C{} }
func (m *PlayerAction_S2C) String() string { return proto.CompactTextString(m) }
func (*PlayerAction_S2C) ProtoMessage()    {}
func (*PlayerAction_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{12}
}

func (m *PlayerAction_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerAction_S2C.Unmarshal(m, b)
}
func (m *PlayerAction_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerAction_S2C.Marshal(b, m, deterministic)
}
func (m *PlayerAction_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAction_S2C.Merge(m, src)
}
func (m *PlayerAction_S2C) XXX_Size() int {
	return xxx_messageInfo_PlayerAction_S2C.Size(m)
}
func (m *PlayerAction_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAction_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAction_S2C proto.InternalMessageInfo

func (m *PlayerAction_S2C) GetRoomData() *RoomData {
	if m != nil {
		return m.RoomData
	}
	return nil
}

type PlayerData struct {
	LoginData            *LoginData `protobuf:"bytes,1,opt,name=loginData,proto3" json:"loginData,omitempty"`
	ContinueVot          float64    `protobuf:"fixed64,2,opt,name=continueVot,proto3" json:"continueVot,omitempty"`
	IsGodGambling        bool       `protobuf:"varint,3,opt,name=IsGodGambling,proto3" json:"IsGodGambling,omitempty"`
	WinCount             int32      `protobuf:"varint,4,opt,name=winCount,proto3" json:"winCount,omitempty"`
	ResultWinMoney       float64    `protobuf:"fixed64,5,opt,name=resultWinMoney,proto3" json:"resultWinMoney,omitempty"`
	ResultLoseMoney      float64    `protobuf:"fixed64,6,opt,name=resultLoseMoney,proto3" json:"resultLoseMoney,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PlayerData) Reset()         { *m = PlayerData{} }
func (m *PlayerData) String() string { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()    {}
func (*PlayerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{13}
}

func (m *PlayerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerData.Unmarshal(m, b)
}
func (m *PlayerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerData.Marshal(b, m, deterministic)
}
func (m *PlayerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerData.Merge(m, src)
}
func (m *PlayerData) XXX_Size() int {
	return xxx_messageInfo_PlayerData.Size(m)
}
func (m *PlayerData) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerData.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerData proto.InternalMessageInfo

func (m *PlayerData) GetLoginData() *LoginData {
	if m != nil {
		return m.LoginData
	}
	return nil
}

func (m *PlayerData) GetContinueVot() float64 {
	if m != nil {
		return m.ContinueVot
	}
	return 0
}

func (m *PlayerData) GetIsGodGambling() bool {
	if m != nil {
		return m.IsGodGambling
	}
	return false
}

func (m *PlayerData) GetWinCount() int32 {
	if m != nil {
		return m.WinCount
	}
	return 0
}

func (m *PlayerData) GetResultWinMoney() float64 {
	if m != nil {
		return m.ResultWinMoney
	}
	return 0
}

func (m *PlayerData) GetResultLoseMoney() float64 {
	if m != nil {
		return m.ResultLoseMoney
	}
	return 0
}

//--11
type MaintainList_S2C struct {
	PlayerList           []*PlayerData `protobuf:"bytes,1,rep,name=playerList,proto3" json:"playerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MaintainList_S2C) Reset()         { *m = MaintainList_S2C{} }
func (m *MaintainList_S2C) String() string { return proto.CompactTextString(m) }
func (*MaintainList_S2C) ProtoMessage()    {}
func (*MaintainList_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{14}
}

func (m *MaintainList_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintainList_S2C.Unmarshal(m, b)
}
func (m *MaintainList_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintainList_S2C.Marshal(b, m, deterministic)
}
func (m *MaintainList_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintainList_S2C.Merge(m, src)
}
func (m *MaintainList_S2C) XXX_Size() int {
	return xxx_messageInfo_MaintainList_S2C.Size(m)
}
func (m *MaintainList_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintainList_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_MaintainList_S2C proto.InternalMessageInfo

func (m *MaintainList_S2C) GetPlayerList() []*PlayerData {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

//--12
type OpenCardResult_S2C struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenCardResult_S2C) Reset()         { *m = OpenCardResult_S2C{} }
func (m *OpenCardResult_S2C) String() string { return proto.CompactTextString(m) }
func (*OpenCardResult_S2C) ProtoMessage()    {}
func (*OpenCardResult_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{15}
}

func (m *OpenCardResult_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenCardResult_S2C.Unmarshal(m, b)
}
func (m *OpenCardResult_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenCardResult_S2C.Marshal(b, m, deterministic)
}
func (m *OpenCardResult_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenCardResult_S2C.Merge(m, src)
}
func (m *OpenCardResult_S2C) XXX_Size() int {
	return xxx_messageInfo_OpenCardResult_S2C.Size(m)
}
func (m *OpenCardResult_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenCardResult_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_OpenCardResult_S2C proto.InternalMessageInfo

//--13
type AmountResult_S2C struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AmountResult_S2C) Reset()         { *m = AmountResult_S2C{} }
func (m *AmountResult_S2C) String() string { return proto.CompactTextString(m) }
func (*AmountResult_S2C) ProtoMessage()    {}
func (*AmountResult_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{16}
}

func (m *AmountResult_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmountResult_S2C.Unmarshal(m, b)
}
func (m *AmountResult_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmountResult_S2C.Marshal(b, m, deterministic)
}
func (m *AmountResult_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmountResult_S2C.Merge(m, src)
}
func (m *AmountResult_S2C) XXX_Size() int {
	return xxx_messageInfo_AmountResult_S2C.Size(m)
}
func (m *AmountResult_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_AmountResult_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_AmountResult_S2C proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Ping)(nil), "pb_msg.Ping")
	proto.RegisterType((*Pong)(nil), "pb_msg.Pong")
	proto.RegisterType((*ErrMsg_S2C)(nil), "pb_msg.ErrMsg_S2C")
	proto.RegisterType((*LoginInfo_C2S)(nil), "pb_msg.LoginInfo_C2S")
	proto.RegisterType((*LoginData)(nil), "pb_msg.LoginData")
	proto.RegisterType((*LoginInfo_S2C)(nil), "pb_msg.LoginInfo_S2C")
	proto.RegisterType((*JoinRoom_C2S)(nil), "pb_msg.JoinRoom_C2S")
	proto.RegisterType((*RoomData)(nil), "pb_msg.RoomData")
	proto.RegisterType((*JoinRoom_S2C)(nil), "pb_msg.JoinRoom_S2C")
	proto.RegisterType((*LeaveRoom_C2S)(nil), "pb_msg.LeaveRoom_C2S")
	proto.RegisterType((*LeaveRoom_S2C)(nil), "pb_msg.LeaveRoom_S2C")
	proto.RegisterType((*PlayerAction_C2S)(nil), "pb_msg.PlayerAction_C2S")
	proto.RegisterType((*PlayerAction_S2C)(nil), "pb_msg.PlayerAction_S2C")
	proto.RegisterType((*PlayerData)(nil), "pb_msg.PlayerData")
	proto.RegisterType((*MaintainList_S2C)(nil), "pb_msg.MaintainList_S2C")
	proto.RegisterType((*OpenCardResult_S2C)(nil), "pb_msg.OpenCardResult_S2C")
	proto.RegisterType((*AmountResult_S2C)(nil), "pb_msg.AmountResult_S2C")
}

func init() { proto.RegisterFile("pb_msg.proto", fileDescriptor_bd13c7201b21be60) }

var fileDescriptor_bd13c7201b21be60 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0xfa, 0x23, 0xb4, 0x5f, 0xd7, 0x2d, 0x58, 0xd3, 0x14, 0x21, 0x84, 0x2a, 0x0b, 0x55,
	0x3d, 0xa0, 0x21, 0x15, 0x89, 0x13, 0x87, 0x4d, 0x05, 0x46, 0x50, 0x0b, 0x95, 0xc7, 0xe0, 0x58,
	0xb9, 0x8d, 0x09, 0xd6, 0x6a, 0x3b, 0x72, 0xdc, 0xa1, 0xfd, 0xdd, 0xfc, 0x03, 0xc8, 0x8e, 0x93,
	0x86, 0x22, 0x0e, 0xe3, 0xe6, 0xf7, 0xfa, 0xfc, 0xbe, 0xf7, 0xbd, 0x5a, 0x81, 0xa3, 0x7c, 0xbd,
	0x12, 0x45, 0x76, 0x9e, 0x6b, 0x65, 0x14, 0x0a, 0x4b, 0x84, 0x43, 0xe8, 0x2c, 0xb9, 0xcc, 0xf0,
	0x18, 0x3a, 0x4b, 0x25, 0x33, 0xf4, 0x0c, 0xa0, 0x60, 0xfa, 0x8e, 0xe9, 0x2f, 0x5c, 0xb0, 0x38,
	0x18, 0x05, 0x93, 0x36, 0x69, 0x30, 0xf8, 0x03, 0xc0, 0x3b, 0xad, 0x17, 0x45, 0xb6, 0xba, 0x9e,
	0xce, 0x50, 0x04, 0x6d, 0x51, 0x64, 0x4e, 0xd6, 0x27, 0xf6, 0x88, 0x4e, 0xa1, 0xcb, 0xb4, 0x56,
	0x3a, 0x6e, 0x39, 0xae, 0x04, 0x08, 0x41, 0x27, 0xa5, 0x86, 0xc6, 0x6d, 0x47, 0xba, 0x33, 0xbe,
	0x81, 0xe1, 0x5c, 0x65, 0x5c, 0x26, 0xf2, 0xbb, 0x5a, 0xcd, 0xa6, 0xd7, 0xe8, 0x18, 0x5a, 0x49,
	0xea, 0xbd, 0x5a, 0x49, 0x8a, 0xce, 0x20, 0xcc, 0xa8, 0x60, 0x49, 0xea, 0xbd, 0x3c, 0x42, 0x4f,
	0xa1, 0x5f, 0x06, 0xba, 0xd1, 0x5b, 0xef, 0xb8, 0x27, 0xf0, 0x2d, 0xf4, 0x9d, 0xed, 0x5b, 0x6a,
	0xe8, 0x5f, 0x96, 0x4f, 0xa0, 0x27, 0xf9, 0xe6, 0xf6, 0x13, 0x15, 0xcc, 0x9b, 0xd6, 0x18, 0xc5,
	0xf0, 0xe8, 0x07, 0xa3, 0x69, 0x22, 0x32, 0x6f, 0x5a, 0x41, 0xfb, 0x0b, 0xdd, 0x6c, 0xd4, 0x4e,
	0x9a, 0xb8, 0x33, 0x0a, 0x26, 0x01, 0xa9, 0x20, 0xbe, 0x68, 0xee, 0x60, 0x0b, 0x79, 0x09, 0xfd,
	0x6d, 0x35, 0xdd, 0xcd, 0x1d, 0x4c, 0x1f, 0x9f, 0xfb, 0xe2, 0xeb, 0x58, 0x64, 0xaf, 0xc1, 0x63,
	0x38, 0xfa, 0xa8, 0xb8, 0x24, 0x4a, 0x09, 0x57, 0xc2, 0x19, 0x84, 0x5a, 0x29, 0x51, 0xa7, 0xf6,
	0x08, 0x63, 0xe8, 0x59, 0x8d, 0xdb, 0xea, 0x5f, 0x9a, 0x37, 0x0d, 0x2f, 0x1b, 0xe6, 0x05, 0xf4,
	0xb4, 0xbf, 0xe3, 0xb3, 0x44, 0x55, 0x96, 0xca, 0x8b, 0xd4, 0x0a, 0x7c, 0x02, 0xc3, 0x39, 0xa3,
	0x77, 0xac, 0x8a, 0xe2, 0x96, 0xab, 0x89, 0xff, 0x5a, 0xee, 0x35, 0x44, 0xcb, 0x2d, 0xbd, 0x67,
	0xfa, 0x72, 0x63, 0xb8, 0x92, 0x6e, 0xc1, 0x08, 0xda, 0x6b, 0x66, 0xdc, 0xf5, 0x80, 0xd8, 0xa3,
	0x65, 0x72, 0x65, 0xdc, 0xff, 0xd1, 0x25, 0xf6, 0x88, 0x2f, 0x0e, 0xee, 0x3d, 0x7c, 0x99, 0x5f,
	0x01, 0x40, 0x69, 0xe1, 0x1a, 0x7b, 0x68, 0x72, 0x34, 0x82, 0xc1, 0x46, 0x49, 0xc3, 0xe5, 0x8e,
	0x7d, 0xf5, 0xd9, 0x02, 0xd2, 0xa4, 0xd0, 0x73, 0x18, 0x26, 0xc5, 0x95, 0x4a, 0xaf, 0xa8, 0x58,
	0x6f, 0xb9, 0x2c, 0x1f, 0x4d, 0x8f, 0xfc, 0x49, 0xda, 0x07, 0xf7, 0x93, 0xcb, 0x59, 0xfd, 0x76,
	0xba, 0xa4, 0xc6, 0x68, 0x0c, 0xc7, 0x9a, 0x15, 0xbb, 0xad, 0xf9, 0xc6, 0xe5, 0x42, 0x49, 0x76,
	0x1f, 0x77, 0xdd, 0x98, 0x03, 0x16, 0x4d, 0xe0, 0xa4, 0x64, 0xe6, 0xaa, 0x60, 0xa5, 0x30, 0x74,
	0xc2, 0x43, 0x1a, 0xbf, 0x87, 0x68, 0x41, 0xb9, 0x34, 0x94, 0xcb, 0x39, 0x2f, 0x8c, 0xeb, 0x6d,
	0x0a, 0x90, 0xbb, 0x22, 0x2c, 0x13, 0x07, 0xa3, 0xf6, 0x64, 0x30, 0x45, 0xd5, 0xee, 0xfb, 0x8a,
	0x48, 0x43, 0x85, 0x4f, 0x01, 0x7d, 0xce, 0x99, 0x9c, 0x51, 0x9d, 0x12, 0x37, 0xc2, 0x3a, 0x61,
	0x04, 0xd1, 0xa5, 0xb0, 0xc9, 0xf7, 0xdc, 0x3a, 0x74, 0x5f, 0x93, 0x57, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x63, 0xb4, 0x66, 0x6e, 0x5d, 0x04, 0x00, 0x00,
}
