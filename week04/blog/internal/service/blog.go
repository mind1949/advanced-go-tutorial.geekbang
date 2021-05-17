package service

import (
	"context"

	pb "blog/api/v1"
	"blog/internal/biz"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type blogService struct {
	articleUc *biz.ArticleUsecase
	*pb.UnimplementedBlogServer
}

func NewBlogService(article *biz.ArticleUsecase) pb.BlogServer {
	return &blogService{articleUc: article}
}

func (b *blogService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.Article, error) {
	a, err := b.articleUc.Create(ctx, req.Title, req.Content)
	if err != nil {
		return nil, err
	}
	return &pb.Article{
		Id:          a.ID,
		Title:       a.Title,
		Content:     a.Content,
		CreatedTime: timestamppb.New(a.CreatedTime),
		UpdatedTime: timestamppb.New(a.UpdatedTime),
	}, nil
}

func (b *blogService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.Article, error) {
	// TODO:
	return nil, nil
}

func (b *blogService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, b.articleUc.Delete(ctx, req.GetId())
}

func (b *blogService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.Article, error) {
	a, err := b.articleUc.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.Article{
		Id:          a.ID,
		Title:       a.Title,
		Content:     a.Content,
		CreatedTime: timestamppb.New(a.CreatedTime),
		UpdatedTime: timestamppb.New(a.UpdatedTime),
	}, nil
}

func (s *blogService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	return &pb.ListArticleReply{Results: []*pb.Article{}}, nil
}
