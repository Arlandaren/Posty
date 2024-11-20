package service

import (
	"context"
	"log"
	"time"

	models "service/internal/shared/storage/dto"

	"service/internal/repository"
	"service/pkg/grpc/posty_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	log.Println("NewService")
	return &Service{
		repo: repo,
	}
}

// Создание нового поста
func (s *Service) CreatePost(ctx context.Context, req *posty_v1.CreatePostRequest, creatorID int64) (*posty_v1.Post, error) {
	postModel := &models.Post{
		Title:     req.Title,
		Image:     req.Image,
		CreatorID: creatorID,
		CreatedAt: time.Now(),
	}

	err := s.repo.CreatePost(postModel)
	if err != nil {
		return nil, err
	}

	// Конвертация модели в protobuf сообщение
	postProto := &posty_v1.Post{
		Id:        postModel.ID,
		Title:     postModel.Title,
		Image:     postModel.Image,
		CreatorId: postModel.CreatorID,
		CreatedAt: timestamppb.New(postModel.CreatedAt),
	}

	return postProto, nil
}

func (s *Service) GetPost(ctx context.Context, id int64) (*posty_v1.Post, error) {
	postModel, err := s.repo.GetPost(id)
	if err != nil {
		return nil, err
	}

	// Конвертация модели в protobuf сообщение
	postProto := &posty_v1.Post{
		Id:        postModel.ID,
		Title:     postModel.Title,
		Image:     postModel.Image,
		CreatorId: postModel.CreatorID,
		CreatedAt: timestamppb.New(postModel.CreatedAt),
	}

	return postProto, nil
}

func (s *Service) GetAllPosts(ctx context.Context, cursor int64, limit int64) ([]*posty_v1.Post, error) {
	postModels, err := s.repo.GetAllPosts(cursor, limit)
	if err != nil {
		return nil, err
	}

	var posts []*posty_v1.Post
	for _, postModel := range postModels {
		postProto := &posty_v1.Post{
			Id:        postModel.ID,
			Title:     postModel.Title,
			Image:     postModel.Image,
			CreatorId: postModel.CreatorID,
			CreatedAt: timestamppb.New(postModel.CreatedAt),
		}
		posts = append(posts, postProto)
	}

	return posts, nil
}

func (s *Service) CommentPost(ctx context.Context, req *posty_v1.CommentPostRequest, creatorID int64) (*posty_v1.Comment, error) {
	commentModel := &models.Comment{
		PostID:    req.PostId,
		Comment:   req.Comment,
		CreatorID: creatorID,
		CreatedAt: time.Now(),
	}

	err := s.repo.CommentPost(commentModel)
	if err != nil {
		return nil, err
	}

	// Конвертация модели в protobuf сообщение
	commentProto := &posty_v1.Comment{
		Id:        commentModel.ID,
		PostId:    commentModel.PostID,
		Comment:   commentModel.Comment,
		CreatorId: commentModel.CreatorID,
		CreatedAt: timestamppb.New(commentModel.CreatedAt),
	}

	return commentProto, nil
}

func (s *Service) GetAllComments(ctx context.Context, postId int64, cursor int64, limit int64) ([]*posty_v1.Comment, error) {
	commentModels, err := s.repo.GetAllComments(postId, cursor, limit)
	if err != nil {
		return nil, err
	}

	var comments []*posty_v1.Comment
	for _, commentModel := range commentModels {
		commentProto := &posty_v1.Comment{
			Id:        commentModel.ID,
			PostId:    commentModel.PostID,
			Comment:   commentModel.Comment,
			CreatorId: commentModel.CreatorID,
			CreatedAt: timestamppb.New(commentModel.CreatedAt),
		}
		comments = append(comments, commentProto)
	}

	return comments, nil
}
