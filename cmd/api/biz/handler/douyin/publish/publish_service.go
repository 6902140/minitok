// Code generated by hertz generator.

package publish

import (
	"bufio"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	publish "minitok/cmd/api/biz/model/douyin/publish"
	"minitok/cmd/api/rpc"
	rpcpublish "minitok/kitex_gen/douyin/publish"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.ActionRequest
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	c.String(consts.StatusBadRequest, err.Error())
	//	return
	//}

	req.Token = c.PostForm("token")
	req.Title = c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	src, err := file.Open()
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	defer src.Close()

	req.Data = make([]byte, file.Size)
	r := bufio.NewReader(src)
	if _, err := r.Read(req.Data); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	resp := new(publish.ActionResponse)
	rpcResp := new(rpcpublish.ActionResponse)

	rpcResp, err = rpc.PublishAction(ctx, &rpcpublish.ActionRequest{
		Title: req.Title,
		Token: req.Token,
		Data:  req.Data,
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

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.ListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.ListResponse)
	rpcResp := new(rpcpublish.ListResponse)

	rpcResp, err = rpc.PublishList(ctx, &rpcpublish.ListRequest{
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
	for i, video := range rpcResp.VideoList {
		if err = copier.Copy(resp.VideoList[i], video); err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
		resp.VideoList[i].ID = video.Id
		resp.VideoList[i].PlayURL = video.PlayUrl
		resp.VideoList[i].CoverURL = video.CoverUrl
		resp.VideoList[i].Author.ID = video.Author.Id
	}
	c.JSON(consts.StatusOK, resp)
}