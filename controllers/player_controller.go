package controllers

import (
	"fmt"
	"strconv"

	"github.com/CatBloom/MahjongMasterApi/service"
	"github.com/gin-gonic/gin"
)

type PlayerController interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type playerController struct {
	s service.PlayerService
}

func NewPlayerController(s service.PlayerService) PlayerController {
	return &playerController{s}
}

func (pc playerController) List(c *gin.Context) {
	lid := c.Params.ByName("lid")
	p, err := pc.s.GetPlayerList(lid)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc playerController) Create(c *gin.Context) {
	p, err := pc.s.CreatePlayer(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc playerController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := pc.s.UpdatePlayer(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc playerController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := pc.s.DeletePlayer(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		//削除成功時にidを返却する
		resId, _ := strconv.Atoi(id)
		c.JSON(200, gin.H{"id": resId})
	}
}
