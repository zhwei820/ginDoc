package sale

import (
	"github.com/gin-gonic/gin"
	"ginDoc/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
)

type ArticleController struct {
}

// @Summary Post
// @Description Post
// @Accept  json
// @Produce  json
// @tag article
// @Param   pets body @Article true "pets fields"
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
// @Param   page query string false  "page"
// @Param   limit query string false  "limit"
// @Success 200 {object} @Articles  "文章列表"
// @Router /article [get]
func (ctrl ArticleController) Get(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	articles := []models.Article{}

	err := db.C(models.CollectionArticle).Find(nil).All(&articles)
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
