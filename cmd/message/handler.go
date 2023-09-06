package main

import (
	"context"
	"minitok/cmd/message/pack"
	"minitok/cmd/message/service"
	message "minitok/kitex_gen/douyin/message"
	"minitok/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.ChatRequest) (*message.ChatResponse, error) {
	if len(req.Token) == 0 || req.ToUserId == 0 {
		return nil, errno.ParamErr
	}
	messageList, err := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		return pack.BuildChatResp(nil, err), nil
	}
	return pack.BuildChatResp(messageList, nil), nil
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.ActionRequest) (*message.ActionResponse, error) {
	if len(req.Token) == 0 || req.ToUserId == 0 {
		return nil, errno.ParamErr
	}
	if err := service.NewMessageActionService(ctx).MessageAction(req); err != nil {
		return pack.BuildActionResp(err), nil
	}
	return pack.BuildActionResp(nil), nil
}
