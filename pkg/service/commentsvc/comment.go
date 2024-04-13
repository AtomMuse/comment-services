package commentsvc

import "atommuse/backend/comment-services/pkg/repositorty/commentrepo"

type CommentService struct {
	Repository commentrepo.CommentRepository
}

func NewCommentService(repo commentrepo.CommentRepository) *CommentService {
	return &CommentService{Repository: repo}
}
