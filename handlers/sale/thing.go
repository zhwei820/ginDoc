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
// @Success 200 {object} @Things  "petslist"
// @Router /things [get]
func (ctrl ThingContoller) List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	things := []models.Thing{}

	err := db.C(models.CollectionStuff).Find(nil).All(&things)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg":fmt.Sprintf("%s", err)})
		return
	}

	// gin.H is a shortcut for map[string]interface{}
	// c.JSON(http.StatusOK, gin.H{"stuff/list": things})
	c.JSON(http.StatusOK, things)
}


// @Summary 新增
// @Description 新增
// @Accept  json
// @Produce  json
// @tag users
// @Param   pets body @Thing true "pets fields"
// @Success 200 {object} @Thing  "success"
// @Router /things [post]
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
// @Success 200 {object} @Pets  "doc"
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
// @Success 200 {object} @Pets  "doc"
// @Router /things [delete]
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
// @Param   pets body @Thing true "pets fields"
// @Success 200 {object} @Thing  "success"
// @Router /things [put]
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
	router.GET("/things/:_id", ctrl.GetOne)
	router.GET("/things", ctrl.List)
	router.POST("/things", ctrl.Create)
	router.DELETE("/things/:_id", ctrl.Delete)
	router.PUT("/things/:_id", ctrl.Put)
}