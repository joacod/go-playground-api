package controllers

import (
	"net/http"
	"playground/models"

	"github.com/gin-gonic/gin"
)

// GetMembers godoc
// @Summary Get all Crew Members
// @Description Get all Crew Members
// @Tags startrek
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /startrek/crewmember [get]
func (ctr *Controller) GetMembers(c *gin.Context) {
	var todo []models.CrewMember
	err := models.GetAllMembers(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// CreateMember godoc
// @Summary Create a Crew Member
// @Description Create a Crew Member
// @Tags startrek
// @Accept json
// @Produce json
// @Param crewMember body models.CrewMember true "Crew member"
// @Success 200 {string} string "ok"
// @Router /startrek/crewmember [post]
func (ctr *Controller) CreateMember(c *gin.Context) {
	var todo models.CrewMember
	c.BindJSON(&todo)
	err := models.CreateMember(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetMember godoc
// @Summary Get a Crew Member
// @Description Get a Crew Member
// @Tags startrek
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.CrewMember
// @Router /startrek/crewmember/{id} [get]
func (ctr *Controller) GetMember(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.CrewMember
	err := models.GetMember(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// UpdateMember godoc
// @Summary Update a Crew Member
// @Description Update a Crew Member
// @Tags startrek
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param crewMember body models.CrewMember true "Crew member"
// @Success 200 {object} models.CrewMember
// @Router /startrek/crewmember/{id} [put]
func (ctr *Controller) UpdateMember(c *gin.Context) {
	var todo models.CrewMember
	id := c.Params.ByName("id")
	err := models.GetMember(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = models.UpdateMember(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// DeleteMember godoc
// @Summary Delete a Crew Member
// @Description Delete a Crew Member
// @Tags startrek
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {string} string "ok"
// @Router /startrek/crewmember/{id} [delete]
func (ctr *Controller) DeleteMember(c *gin.Context) {
	var todo models.CrewMember
	id := c.Params.ByName("id")
	err := models.DeleteMember(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
