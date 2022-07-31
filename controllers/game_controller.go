package controllers

import (
	"fmt"

	"github.com/CatBloom/MahjongMasterApi/service"
	"github.com/gin-gonic/gin"
)

type GameController interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type gameController struct {
	s service.GameService
}

func NewGameController(s service.GameService) GameController {
	return &gameController{s}
}

func (gc gameController) List(c *gin.Context) {
	lid := c.Params.ByName("lid")
	g, err := gc.s.GetGameList(lid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		fmt.Println(g)
		c.JSON(200, g)
	}
}

func (gc gameController) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	l, err := gc.s.GetGame(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		fmt.Println(l)
		c.JSON(200, l)
	}
}

func (gc gameController) Create(c *gin.Context) {
	l, err := gc.s.CreateGame(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, l)
	}
}

func (gc gameController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	l, err := gc.s.UpdateGame(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, l)
	}
}

func (gc gameController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := gc.s.DeleteGame(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}
