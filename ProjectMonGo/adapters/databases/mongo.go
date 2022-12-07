package databases

import (
	"ProjectMonGo/adapters/databases/mapper"
	"ProjectMonGo/api/Request"
	"ProjectMonGo/core/entities"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PostServiceImpl struct {
	postCollection *mongo.Collection
	ctx            context.Context
}

func NewMonGoDb(collection *mongo.Collection) *PostServiceImpl {
	return &PostServiceImpl{
		postCollection: collection,
	}
}

func (p *PostServiceImpl) CreatePost(post *Request.CreatePostRequest) (*entities.DBPost, error)  {
	post.CreateAt = time.Now()
	post.UpdatedAt = post.CreateAt

	userModel := &Request.CreatePostRequest{
		Title:     post.Title,
		Content:   post.Content,
		Image:     post.Image,
		User:      post.User,
		CreateAt:  time.Time{},
		UpdatedAt: time.Time{},
	}
	_, err := p.postCollection.InsertOne(p.ctx, userModel)
	if err != nil {
		return nil,err
	}
	return nil,nil
}
///UpdatePost(string,*Request.UpdatePost)(*Request.UpdatePost,error)
func (p *PostServiceImpl) UpdatePost(id string, data *Request.UpdatePost) (*Request.UpdatePost, error) {
	doc, err := mapper.ToDoc(data)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "title", Value: id}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.postCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatePost *Request.UpdatePost

	if err := res.Decode(&updatePost); err != nil {
		return nil, errors.New("no post with that Id exists")
	}

	return updatePost, nil
}
func (p *PostServiceImpl) FindPostById(id string) (*entities.DBPost, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}
	var post *entities.DBPost
	if err := p.postCollection.FindOne(p.ctx, query).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}
		return nil, err
	}
	return post, nil
}
func (p *PostServiceImpl) FindPosts(page int, limit int) ([]*entities.DBPost, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	skip := (page - 1) * limit
	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))

	query := bson.M{}
	cursor,err := p.postCollection.Find(p.ctx,query,&opt)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(p.ctx)

	var posts []*entities.DBPost

	for cursor.Next(p.ctx) {
		post := &entities.DBPost{}
		err := cursor.Decode(post)
		if err != nil {
			return nil, err
		}
		posts=append(posts,post)
	}
	if err:= cursor.Err();err != nil{
		return nil,err
	}
	if len(posts)==0{
		return []*entities.DBPost{},nil
	}
	return posts, nil
}
func (p *PostServiceImpl) DeletePosts(title string) error {

	filter := bson.M{"title":title}
	_, err := p.postCollection.DeleteOne(p.ctx, filter)
	if err != nil {
		return err
	}
	return nil
}















