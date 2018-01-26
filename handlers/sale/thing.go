package sale

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ginDoc/models"
	"github.com/gin-gonic/gin"
	"fmt"
)

type ThingContoller struct {

}

// @Summary 列表
// @Description 列表
// @Accept  html
// @Produce  json
// @tag users
// @Param   page query string false  "page"
// @Success 200 {object} @Things  "Things"
// @Router /thing [get]
func (ctrl ThingContoller) List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	thing := []models.Thing{}

	err := db.C(models.CollectionStuff).Find(nil).All(&thing)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg":fmt.Sprintf("%s", err)})
		return
	}

	// gin.H is a shortcut for map[string]interface{}
	// c.JSON(http.StatusOK, gin.H{"stuff/list": thing})
	c.JSON(http.StatusOK, thing)
}


// @Summary 新增
// @Description 新增
// @Accept  json
// @Produce  json
// @tag users
// @Param   thing body @Thing true "thing"
// @Success 200 {object} @Thing  "success"
// @Router /thing [post]
func (ctrl ThingContoller) Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	err := c.ShouldBindJSON(&thing)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionStuff).Insert(thing)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, thing)
}


// @Summary 单个
// @Description 单个
// @Accept  json
// @Produce  json
// @tag users
// @Param   _id query string false  "_id of doc"
// @Success 200 {object} @Thing  "doc"
// @Router /thing [get]
func (ctrl ThingContoller) GetOne(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionStuff).FindId(oID).One(&thing)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, thing)
}

// @Summary Delete
// @Description Delete
// @Accept  html
// @Produce  json
// @tag users
// @Param   _id query string false  "_id of doc"
// @Success 200 {object} @Thing  "doc"
// @Router /thing [delete]
func (ctrl ThingContoller) Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionStuff).RemoveId(oID)
	if err != nil {
		c.Error(err)
	}

	// What to do here if this is close to REST
	//c.Redirect(http.StatusMovedPermanently, "/stuff")
	c.Data(204, "application/json", make([]byte, 0))
}

// @Summary Put
// @Description Put
// @Accept  json
// @Produce  json
// @tag users
// @Param   thing body @Thing true "thing fields"
// @Param   _id query string true "mongo id"
// @Success 200 {object} @Thing  "success"
// @Router /thing [put]
func (ctrl ThingContoller) Put(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	err := c.Bind(&thing)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{ "_id": bson.ObjectIdHex(c.Param("_id")) }
	doc := bson.M{
		"name":		thing.Name,
		"value":	thing.Value,
	}
	err = db.C(models.CollectionStuff).Update(query, doc)
	if err != nil {
		c.Error(err)
	}

	c.Data(http.StatusOK, "application/json", make([]byte, 0))
	//c.Rediret(http.StatusMovedPermanently, "/stuff"
}


func (ctrl ThingContoller) SetRoute(router *gin.Engine) {
	router.GET("/thing/:_id", ctrl.GetOne)
	router.GET("/thing", ctrl.List)
	router.POST("/thing", ctrl.Create)
	router.DELETE("/thing/:_id", ctrl.Delete)
	router.PUT("/thing/:_id", ctrl.Put)
}