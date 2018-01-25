package things

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ginDoc/models"
	"github.com/gin-gonic/gin"
	"fmt"
)


// @Summary getPets
// @Description 获取pets
// @ID file.upload
// @Accept  json
// @Produce  json
// @tag users
// @Param   page query string false  "page of the gets"
// @Success 200 {object} @Things  "petslist"
// @Router /things [get]
func List(c *gin.Context) {
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


// @Summary getPets
// @Description 创建thing
// @ID create_thing
// @Accept  json
// @Produce  json
// @tag users
// @Param   pets body @Thing true "pets fields"
// @Success 200 {object} @Thing  "success"
// @Router /things [post]
func Create(c *gin.Context) {
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


// @Summary getOne
// @Description getOne
// @ID getOne
// @Accept  json
// @Produce  json
// @tag users
// @Param   _id query string false  "_id of doc"
// @Success 200 {object} @Pets  "doc"
// @Router /thing [get]
func GetOne(c *gin.Context) {
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
// @ID Delete
// @Accept  html
// @Produce  json
// @tag users
// @Param   _id query string false  "_id of doc"
// @Success 200 {object} @Pets  "doc"
// @Router /things [delete]
func Delete(c *gin.Context) {
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

// @Summary getPets
// @Description 创建thing
// @ID create_thing
// @Accept  json
// @Produce  json
// @tag users
// @Param   pets body @Thing true "pets fields"
// @Success 200 {object} @Thing  "success"
// @Router /things [put]
func Update(c *gin.Context) {
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
