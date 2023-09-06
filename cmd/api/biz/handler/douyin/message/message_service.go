// Code generated by hertz generator.

package message

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	message "minitok/cmd/api/biz/model/douyin/message"
	"minitok/cmd/api/rpc"
	rpcmessage "minitok/kitex_gen/douyin/message"
)

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.ChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(message.ChatResponse)
	rpcResp := new(rpcmessage.ChatResponse)

	fmt.Printf("%#v\n", req)

	rpcResp, err = rpc.MessageChat(ctx, &rpcmessage.ChatRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		PreMsgTime: req.PreMsgTime,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	if err = copier.Copy(resp, rpcResp); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	for i, msg := range rpcResp.MessageList {
		if err = copier.Copy(resp.MessageList[i], msg); err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
		resp.MessageList[i].ID = msg.Id
		resp.MessageList[i].FromUserID = msg.FromUserId
		resp.MessageList[i].ToUserID = msg.ToUserId
	}

	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.ActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(message.ActionResponse)
	rpcResp := new(rpcmessage.ActionResponse)

	rpcResp, err = rpc.MessageAction(ctx, &rpcmessage.ActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
		Content:    req.Content,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	if err = copier.Copy(resp, rpcResp); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}
