package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	newUser, err := s.uu.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Profile: &v1.Profile{
			Username:  newUser.Username,
			Bio:       newUser.Bio,
			Image:     newUser.Image,
			Following: false,
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
