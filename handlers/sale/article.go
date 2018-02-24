package sale

import (
	"github.com/gin-gonic/gin"
	"ginDoc/models"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
	"strconv"
	"ginDoc/db"
	clog "github.com/cihub/seelog"
	"github.com/garyburd/redigo/redis"
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

	c.JSON(http.StatusOK, gin.H{"data": articles})
}

// @Summary 压测
// @Description 压测
// @Produce  json
// @tag article
// @Success 200 {object} @Articles  "压测下"
// @Router /siege [get]
func (ctrl ArticleController) boom2(c *gin.Context) {
	rc := db.CachePool.Get()
	defer rc.Close()
	key1 := "s"
	rc.Do("SET", key1, 0)
	n, err := redis.String(rc.Do("GET", key1))
	if err != nil {
		clog.Errorf("redis %s: %s", key1, err.Error())
	}
	key2 := "hset"
	rc.Do("HSET", key2, "a", "ddd")
	a, err := redis.String(rc.Do("HGET", key2, "a"))
	if err != nil {
		clog.Errorf("redis %s:%s: %s", key2, "a", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"n": n, "a": a}})
}

// @Summary 压测redis get
// @Description 压测
// @Produce  json
// @tag article
// @Success 200 {object} @Articles  "压测下"
// @Router /siege [get]
func (ctrl ArticleController) boom1(c *gin.Context) {
	rc := db.CachePool.Get()
	defer rc.Close()
	key1 := "s"
	n, err := redis.String(rc.Do("GET", key1))
	if err != nil {
		clog.Errorf("redis %s: %s", key1, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"n": n,}})
}

// @Summary 压测 hello
// @Description 压测
// @Produce  json
// @tag article
// @Success 200 {object} @Articles  "压测下"
// @Router /siege [get]
func (ctrl ArticleController) hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"n": 11,}})
}

func (ctrl ArticleController) SetRoute(router *gin.Engine) {
	router.GET("/article", ctrl.Get)
	router.POST("/article", ctrl.Post)
	router.GET("/boom2", ctrl.boom2)
	router.GET("/boom1", ctrl.boom1)
	router.GET("/hello", ctrl.hello)

}
