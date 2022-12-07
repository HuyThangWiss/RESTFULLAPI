package posts

import (
	"ProjectMonGo/adapters/databases"
	"ProjectMonGo/api/Request"
	"ProjectMonGo/core/entities"
)

type PostService interface {
	CreatePost(request *Request.CreatePostRequest)(*entities.DBPost,error)
	UpdatePost(string,*Request.UpdatePost)(*Request.UpdatePost,error)
	FindPostById(string)(*entities.DBPost,error)
	FindPosts(page int, limit int)([]*entities.DBPost,error)
	DeletePosts(string) error
}

func InitUserRepositoryPort(mongoDb *databases.PostServiceImpl) PostService{
	return mongoDb
}


