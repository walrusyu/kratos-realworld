package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.LoginReply, err error) {
	return &v1.LoginReply{
		Profile: &v1.Profile{
			Username:  "testUser",
			Bio:       "testBio",
			Image:     "testImg",
			Following: true,
		},
	}, nil
}

func (s *RealWorldService) FollowUser(context.Context, *v1.FollowUserRequest) (*v1.FollowUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (s *RealWorldService) UnFollowUser(context.Context, *v1.UnFollowUserRequest) (*v1.UnFollowUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (s *RealWorldService) ListArticles(context.Context, *v1.ListArticlesRequest) (*v1.ListArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (s *RealWorldService) FeedArticles(context.Context, *v1.FeedArticlesRequest) (*v1.FeedArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (s *RealWorldService) GetArticle(context.Context, *v1.GetArticleRequest) (*v1.GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (s *RealWorldService) CreateArticle(context.Context, *v1.CreateArticleRequest) (*v1.CreateArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (s *RealWorldService) UpdateArticle(context.Context, *v1.UpdateArticleRequest) (*v1.UpdateArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (s *RealWorldService) DeleteArticle(context.Context, *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (s *RealWorldService) AddComment(context.Context, *v1.AddCommentRequest) (*v1.AddCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (s *RealWorldService) GetComments(context.Context, *v1.GetCommentsRequest) (*v1.GetCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (s *RealWorldService) DeleteComment(context.Context, *v1.DeleteCommentRequest) (*v1.DeleteCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (s *RealWorldService) FavoriteArticle(context.Context, *v1.FavoriteArticleRequest) (*v1.FavoriteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}
func (s *RealWorldService) UnFavoriteArticle(context.Context, *v1.UnFavoriteArticleRequest) (*v1.UnFavoriteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavoriteArticle not implemented")
}
func (s *RealWorldService) GetTags(context.Context, *v1.GetTagsRequest) (*v1.GetTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
