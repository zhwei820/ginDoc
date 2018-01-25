package handlers

import (
	"github.com/gin-gonic/gin"
	"ginDoc/models"
	"fmt"
	"gopkg.in/mgo.v2"
)

//ArticleController ...
type ArticleController struct{}

//Create ...
func (ctrl ArticleController) Post(c *gin.Context) {

	db := c.MustGet("db").(*mgo.Database)

	article := models.Article{}
	err := c.ShouldBindJSON(&article)

	if err != nil {
		c.JSON(406, gin.H{"message": fmt.Sprintf("%s", err)})
		c.Abort()
		return
	}

	err = db.C(models.CollectionStuff).Insert(thing)

	c.JSON(200, gin.H{"message": "Article created", "code": 0})
}

//All ...
func (ctrl ArticleController) Get(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	data, err := articleModel.All(userID)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the articles", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}
