package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
	myerrors "kratos-realworld/internal/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	if len(req.Email) == 0 {
		return nil, myerrors.NewError(411, "email", "can not be empty")
	}

	newUser, err := s.uu.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		User: &v1.User{
			Username: newUser.Username,
			Email:    newUser.Email,
			Bio:      newUser.Bio,
			Image:    newUser.Image,
			Token:    newUser.Token,
		},
	}, nil
}
func (s *RealWorldService) GetCurrentUser(context.Context, *emptypb.Empty) (*v1.GetCurrentUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (s *RealWorldService) UpdateUser(context.Context, *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (s *RealWorldService) GetProfile(context.Context, *v1.GetProfileRequest) (*v1.GetProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
