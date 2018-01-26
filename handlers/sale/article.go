package sale

import (
	"github.com/gin-gonic/gin"
	"ginDoc/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
	"strconv"
)

type ArticleController struct {
}

// @Summary Post
// @Description Post
// @Accept  json
// @Produce  json
// @tag article
// @Param   article body @Article true "article"
// @Success 200 {object} @Article  "success"
// @Router /article [post]
func (ctrl ArticleController) Post(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}
	err := c.ShouldBindJSON(&article)

	if err != nil {
		c.JSON(406, gin.H{"message": fmt.Sprintf("%s", err)})
		c.Abort()
		return
	}

	err = db.C(models.CollectionArticle).Insert(article)

	c.JSON(200, gin.H{"message": "Article created", "code": 0})
}

// @Summary 文章列表
// @Description 文章列表
// @Accept  json
// @Produce  json
// @tag article
// @Param   page query int false  "page"
// @Param   limit query int false  "limit"
// @Success 200 {object} @Articles  "文章列表"
// @Router /article [get]
func (ctrl ArticleController) Get(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var articles []models.Article

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	err = db.C(models.CollectionArticle).Find(nil).
		Skip(limit * (page - 1)).Limit(limit).All(&articles)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("%s", err)})
		return
	}

	c.JSON(200, gin.H{"data": articles})
}

func (ctrl ArticleController) SetRoute(router *gin.Engine) {
	router.GET("/article", ctrl.Get)
	router.POST("/article", ctrl.Post)

}
