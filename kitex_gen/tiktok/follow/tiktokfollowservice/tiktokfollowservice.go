// Code generated by Kitex v0.4.4. DO NOT EDIT.

package tiktokfollowservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	follow "github.com/ozline/tiktok/kitex_gen/tiktok/follow"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return tiktokFollowServiceServiceInfo
}

var tiktokFollowServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "tiktokFollowService"
	handlerType := (*follow.TiktokFollowService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Ping":           kitex.NewMethodInfo(pingHandler, newPingArgs, newPingResult, false),
		"RelationAction": kitex.NewMethodInfo(relationActionHandler, newRelationActionArgs, newRelationActionResult, false),
		"RelationQuery":  kitex.NewMethodInfo(relationQueryHandler, newRelationQueryArgs, newRelationQueryResult, false),
		"FollowList":     kitex.NewMethodInfo(followListHandler, newFollowListArgs, newFollowListResult, false),
		"FollowerList":   kitex.NewMethodInfo(followerListHandler, newFollowerListArgs, newFollowerListResult, false),
		"FriendList":     kitex.NewMethodInfo(friendListHandler, newFriendListArgs, newFriendListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "follow",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func pingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(follow.PingReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(follow.TiktokFollowService).Ping(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PingArgs:
		success, err := handler.(follow.TiktokFollowService).Ping(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PingResult)
		realResult.Success = success
	}
	return nil
}
func newPingArgs() interface{} {
	return &PingArgs{}
}

func newPingResult() interface{} {
	return &PingResult{}
}

type PingArgs struct {
	Req *follow.PingReq
}

func (p *PingArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(follow.PingReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PingArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PingArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PingArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PingArgs) Unmarshal(in []byte) error {
	msg := new(follow.PingReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PingArgs_Req_DEFAULT *follow.PingReq

func (p *PingArgs) GetReq() *follow.PingReq {
	if !p.IsSetReq() {
		return PingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PingArgs) IsSetReq() bool {
	return p.Req != nil
}

type PingResult struct {
	Success *follow.BaseRsp
}

var PingResult_Success_DEFAULT *follow.BaseRsp

func (p *PingResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(follow.BaseRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PingResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PingResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PingResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PingResult) Unmarshal(in []byte) error {
	msg := new(follow.BaseRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PingResult) GetSuccess() *follow.BaseRsp {
	if !p.IsSetSuccess() {
		return PingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PingResult) SetSuccess(x interface{}) {
	p.Success = x.(*follow.BaseRsp)
}

func (p *PingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func relationActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(follow.RelationActionReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(follow.TiktokFollowService).RelationAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RelationActionArgs:
		success, err := handler.(follow.TiktokFollowService).RelationAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RelationActionResult)
		realResult.Success = success
	}
	return nil
}
func newRelationActionArgs() interface{} {
	return &RelationActionArgs{}
}

func newRelationActionResult() interface{} {
	return &RelationActionResult{}
}

type RelationActionArgs struct {
	Req *follow.RelationActionReq
}

func (p *RelationActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(follow.RelationActionReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RelationActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RelationActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RelationActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RelationActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RelationActionArgs) Unmarshal(in []byte) error {
	msg := new(follow.RelationActionReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RelationActionArgs_Req_DEFAULT *follow.RelationActionReq

func (p *RelationActionArgs) GetReq() *follow.RelationActionReq {
	if !p.IsSetReq() {
		return RelationActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RelationActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type RelationActionResult struct {
	Success *follow.BaseRsp
}

var RelationActionResult_Success_DEFAULT *follow.BaseRsp

func (p *RelationActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(follow.BaseRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RelationActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RelationActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RelationActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RelationActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RelationActionResult) Unmarshal(in []byte) error {
	msg := new(follow.BaseRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RelationActionResult) GetSuccess() *follow.BaseRsp {
	if !p.IsSetSuccess() {
		return RelationActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RelationActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*follow.BaseRsp)
}

func (p *RelationActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func relationQueryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(follow.RelationQueryReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(follow.TiktokFollowService).RelationQuery(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RelationQueryArgs:
		success, err := handler.(follow.TiktokFollowService).RelationQuery(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RelationQueryResult)
		realResult.Success = success
	}
	return nil
}
func newRelationQueryArgs() interface{} {
	return &RelationQueryArgs{}
}

func newRelationQueryResult() interface{} {
	return &RelationQueryResult{}
}

type RelationQueryArgs struct {
	Req *follow.RelationQueryReq
}

func (p *RelationQueryArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(follow.RelationQueryReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RelationQueryArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RelationQueryArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RelationQueryArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RelationQueryArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RelationQueryArgs) Unmarshal(in []byte) error {
	msg := new(follow.RelationQueryReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RelationQueryArgs_Req_DEFAULT *follow.RelationQueryReq

func (p *RelationQueryArgs) GetReq() *follow.RelationQueryReq {
	if !p.IsSetReq() {
		return RelationQueryArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RelationQueryArgs) IsSetReq() bool {
	return p.Req != nil
}

type RelationQueryResult struct {
	Success *follow.RelationQueryRsp
}

var RelationQueryResult_Success_DEFAULT *follow.RelationQueryRsp

func (p *RelationQueryResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(follow.RelationQueryRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RelationQueryResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RelationQueryResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RelationQueryResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RelationQueryResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RelationQueryResult) Unmarshal(in []byte) error {
	msg := new(follow.RelationQueryRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RelationQueryResult) GetSuccess() *follow.RelationQueryRsp {
	if !p.IsSetSuccess() {
		return RelationQueryResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RelationQueryResult) SetSuccess(x interface{}) {
	p.Success = x.(*follow.RelationQueryRsp)
}

func (p *RelationQueryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(follow.UserListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(follow.TiktokFollowService).FollowList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowListArgs:
		success, err := handler.(follow.TiktokFollowService).FollowList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowListArgs() interface{} {
	return &FollowListArgs{}
}

func newFollowListResult() interface{} {
	return &FollowListResult{}
}

type FollowListArgs struct {
	Req *follow.UserListReq
}

func (p *FollowListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(follow.UserListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowListArgs) Unmarshal(in []byte) error {
	msg := new(follow.UserListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowListArgs_Req_DEFAULT *follow.UserListReq

func (p *FollowListArgs) GetReq() *follow.UserListReq {
	if !p.IsSetReq() {
		return FollowListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowListResult struct {
	Success *follow.UserListRsp
}

var FollowListResult_Success_DEFAULT *follow.UserListRsp

func (p *FollowListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(follow.UserListRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowListResult) Unmarshal(in []byte) error {
	msg := new(follow.UserListRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowListResult) GetSuccess() *follow.UserListRsp {
	if !p.IsSetSuccess() {
		return FollowListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowListResult) SetSuccess(x interface{}) {
	p.Success = x.(*follow.UserListRsp)
}

func (p *FollowListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(follow.UserListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(follow.TiktokFollowService).FollowerList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowerListArgs:
		success, err := handler.(follow.TiktokFollowService).FollowerList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowerListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowerListArgs() interface{} {
	return &FollowerListArgs{}
}

func newFollowerListResult() interface{} {
	return &FollowerListResult{}
}

type FollowerListArgs struct {
	Req *follow.UserListReq
}

func (p *FollowerListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(follow.UserListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowerListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowerListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowerListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowerListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowerListArgs) Unmarshal(in []byte) error {
	msg := new(follow.UserListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowerListArgs_Req_DEFAULT *follow.UserListReq

func (p *FollowerListArgs) GetReq() *follow.UserListReq {
	if !p.IsSetReq() {
		return FollowerListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowerListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowerListResult struct {
	Success *follow.UserListRsp
}

var FollowerListResult_Success_DEFAULT *follow.UserListRsp

func (p *FollowerListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(follow.UserListRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowerListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowerListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowerListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowerListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowerListResult) Unmarshal(in []byte) error {
	msg := new(follow.UserListRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowerListResult) GetSuccess() *follow.UserListRsp {
	if !p.IsSetSuccess() {
		return FollowerListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowerListResult) SetSuccess(x interface{}) {
	p.Success = x.(*follow.UserListRsp)
}

func (p *FollowerListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(follow.UserListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(follow.TiktokFollowService).FriendList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FriendListArgs:
		success, err := handler.(follow.TiktokFollowService).FriendList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FriendListResult)
		realResult.Success = success
	}
	return nil
}
func newFriendListArgs() interface{} {
	return &FriendListArgs{}
}

func newFriendListResult() interface{} {
	return &FriendListResult{}
}

type FriendListArgs struct {
	Req *follow.UserListReq
}

func (p *FriendListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(follow.UserListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FriendListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FriendListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FriendListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FriendListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FriendListArgs) Unmarshal(in []byte) error {
	msg := new(follow.UserListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FriendListArgs_Req_DEFAULT *follow.UserListReq

func (p *FriendListArgs) GetReq() *follow.UserListReq {
	if !p.IsSetReq() {
		return FriendListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FriendListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FriendListResult struct {
	Success *follow.UserListRsp
}

var FriendListResult_Success_DEFAULT *follow.UserListRsp

func (p *FriendListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(follow.UserListRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FriendListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FriendListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FriendListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FriendListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FriendListResult) Unmarshal(in []byte) error {
	msg := new(follow.UserListRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FriendListResult) GetSuccess() *follow.UserListRsp {
	if !p.IsSetSuccess() {
		return FriendListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FriendListResult) SetSuccess(x interface{}) {
	p.Success = x.(*follow.UserListRsp)
}

func (p *FriendListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Ping(ctx context.Context, Req *follow.PingReq) (r *follow.BaseRsp, err error) {
	var _args PingArgs
	_args.Req = Req
	var _result PingResult
	if err = p.c.Call(ctx, "Ping", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationAction(ctx context.Context, Req *follow.RelationActionReq) (r *follow.BaseRsp, err error) {
	var _args RelationActionArgs
	_args.Req = Req
	var _result RelationActionResult
	if err = p.c.Call(ctx, "RelationAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RelationQuery(ctx context.Context, Req *follow.RelationQueryReq) (r *follow.RelationQueryRsp, err error) {
	var _args RelationQueryArgs
	_args.Req = Req
	var _result RelationQueryResult
	if err = p.c.Call(ctx, "RelationQuery", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, Req *follow.UserListReq) (r *follow.UserListRsp, err error) {
	var _args FollowListArgs
	_args.Req = Req
	var _result FollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, Req *follow.UserListReq) (r *follow.UserListRsp, err error) {
	var _args FollowerListArgs
	_args.Req = Req
	var _result FollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, Req *follow.UserListReq) (r *follow.UserListRsp, err error) {
	var _args FriendListArgs
	_args.Req = Req
	var _result FriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}