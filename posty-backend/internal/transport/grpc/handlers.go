package transport

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"service/internal/service"
	"service/pkg/grpc/posty_v1"
)

type Server struct {
	posty_v1.UnimplementedPostyServiceServer
	Service *service.Service
}

func NewServer(service *service.Service) *Server {
	return &Server{
		Service: service,
	}
}

// CreatePost handler
func (s *Server) CreatePost(ctx context.Context, req *posty_v1.CreatePostRequest) (*posty_v1.Post, error) {
	// Retrieve userID from context
	userIDInterface := ctx.Value("userID")
	if userIDInterface == nil {
		return nil, status.Error(codes.Unauthenticated, "User not authenticated")
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		return nil, status.Error(codes.Internal, "Invalid user ID type")
	}

	// Convert int to int64 if necessary
	userIDInt64 := int64(userID)

	// Call service method
	post, err := s.Service.CreatePost(ctx, req, userIDInt64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create post: %v", err)
	}

	return post, nil
}

// GetPost handler
func (s *Server) GetPost(ctx context.Context, req *posty_v1.GetPostRequest) (*posty_v1.Post, error) {
	post, err := s.Service.GetPost(ctx, req.Id)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, status.Errorf(codes.NotFound, "Post not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get post: %v", err)
	}

	return post, nil
}

// GetAllPosts handler
func (s *Server) GetAllPosts(ctx context.Context, req *posty_v1.GetAllPostsRequest) (*posty_v1.GetAllPostsResponse, error) {
	posts, err := s.Service.GetAllPosts(ctx, req.Cursor, req.Limit)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get posts: %v", err)
	}

	return &posty_v1.GetAllPostsResponse{
		Posts: posts,
	}, nil
}

// CommentPost handler
func (s *Server) CommentPost(ctx context.Context, req *posty_v1.CommentPostRequest) (*posty_v1.Comment, error) {
	userIDInterface := ctx.Value("userID")
	if userIDInterface == nil {
		return nil, status.Error(codes.Unauthenticated, "User not authenticated")
	}
	userID, ok := userIDInterface.(int)
	if !ok {
		return nil, status.Error(codes.Internal, "Invalid user ID type")
	}

	// Convert int to int64 if necessary
	userIDInt64 := int64(userID)

	// Вызов сервисного метода
	comment, err := s.Service.CommentPost(ctx, req, userIDInt64)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to comment on post: %v", err)
	}

	return comment, nil
}

// GetAllComments handler
func (s *Server) GetAllComments(ctx context.Context, req *posty_v1.GetAllCommentsRequest) (*posty_v1.GetAllCommentsResponse, error) {
	comments, err := s.Service.GetAllComments(ctx, req.PostId, req.Cursor, req.Limit)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get comments: %v", err)
	}

	return &posty_v1.GetAllCommentsResponse{
		Comments: comments,
	}, nil
}
