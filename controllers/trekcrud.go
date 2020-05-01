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
	var member []models.CrewMember
	err := models.GetAllMembers(&member)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
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
	var member models.CrewMember
	c.BindJSON(&member)
	err := models.CreateMember(&member)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
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
	var member models.CrewMember
	err := models.GetMember(&member, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
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
	var member models.CrewMember
	id := c.Params.ByName("id")
	err := models.GetMember(&member, id)
	if err != nil {
		c.JSON(http.StatusNotFound, member)
	}
	c.BindJSON(&member)
	err = models.UpdateMember(&member, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
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
	var member models.CrewMember
	id := c.Params.ByName("id")
	err := models.DeleteMember(&member, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}

// CreateSampleCrew godoc
// @Summary Create a Sample Crew
// @Description Create a Sample Crew Member with 6 Classic Trek Characters
// @Tags startrek
// @Accept json
// @Produce json
// @Success 200 {string} string "Classic Star Trek Crew created"
// @Router /startrek/samplecrew [post]
func (ctr *Controller) CreateSampleCrew(c *gin.Context) {
	crew := classicCrew()
	var err error

	for _, member := range crew {
		err = models.CreateMember(&member)
		if err != nil {
			break
		}
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)			
	} else {
		c.JSON(http.StatusOK, crew)
	}	
}

func classicCrew() []models.CrewMember {
	crew := []models.CrewMember{
		{
			Name: "Mr. Spock",
			Starship: "USS Enterprise NCC-1701",
		},
		{
			Name: "James T. Kirk",
			Starship: "USS Enterprise NCC-1701",
		},
		{
			Name: "Leonard McCoy",
			Starship: "USS Enterprise NCC-1701",
		},
		{
			Name: "Jean Luc Picard",
			Starship: "USS Enterprise NCC-1701-D",
		},
		{
			Name: "Data",
			Starship: "USS Enterprise NCC-1701-D",
		},
		{
			Name: "William Riker",
			Starship: "USS Enterprise NCC-1701-D",
		},
	}

	return crew
}
