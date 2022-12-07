package Controller

import (
	"ProjectMonGo/api/Request"
	"ProjectMonGo/core/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)


type PostController struct {
	postServices *Services.PostService
}

func NewPostServiceS(postService *Services.PostService)*PostController  {
	return &PostController{postServices: postService}
}

func (u *PostController)Create(c *gin.Context)  {
	var PostReq Request.CreatePostRequest
	if err := c.ShouldBindJSON(&PostReq);err != nil{
		c.JSON(http.StatusInternalServerError, "lỗi rồi nhé")
		return
	}
	_,err := u.postServices.Create_Post(c,&PostReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK, PostReq)
}

func (u *PostController)FindALl(ctx *gin.Context)  {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	posts,err1 := u.postServices.FindPost(intPage,intLimit)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": posts})
}
func (p *PostController)DeletePost(ctx *gin.Context)  {
	postId := ctx.Param("title")
	err := p.postServices.Delete_Post(postId)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, "success")
}

func (p *PostController)UpdatePost(ctx *gin.Context)  {
	postId := ctx.Param("title")
	var post *Request.UpdatePost
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	updatePost,err := p.postServices.Update_Post(postId,post)
	if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatePost})

}
func (p *PostController)FindPostId(ctx *gin.Context)  {
	postId := ctx.Param("postId")
	post,err := p.postServices.Find_Id(postId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}



