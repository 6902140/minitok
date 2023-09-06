// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	user "minitok/cmd/api/biz/model/douyin/user"
	"minitok/cmd/api/rpc"
	rpcuser "minitok/kitex_gen/douyin/user"
)

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.RegisterResponse)
	rpcResp := new(rpcuser.RegisterResponse)

	rpcResp, err = rpc.UserRegister(ctx, &rpcuser.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	if err = copier.Copy(resp, rpcResp); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp.UserID = rpcResp.UserId
	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.LoginResponse)
	rpcResp := new(rpcuser.LoginResponse)

	rpcResp, err = rpc.UserLogin(ctx, &rpcuser.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	if err = copier.Copy(resp, rpcResp); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp.UserID = rpcResp.UserId

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.InfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.InfoResponse)
	rpcResp := new(rpcuser.InfoResponse)

	rpcResp, err = rpc.UserInfo(ctx, &rpcuser.InfoRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	if err = copier.Copy(resp, rpcResp); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	if err = copier.Copy(resp.User, rpcResp.User); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp.User.ID = rpcResp.User.Id

	c.JSON(consts.StatusOK, resp)
}
