package controllers

import (
	"errors"
	"strings"
	"wackdo/src/initializers"
	"wackdo/src/models"

	"github.com/gin-gonic/gin"
)

type EmployeePostRequest struct {
	Name  string        `json:"name" binding:"required"`
	Roles []models.Role `json:"roles" binding:"required"`
}

func PostEmployee(c *gin.Context) {
	var req EmployeePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ValidateEmployeePostRequest(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	insertEmployee := &models.Employee{
		Name:  req.Name,
		Roles: req.Roles,
	}
	initializers.DB.Create(insertEmployee)

	c.JSON(200, gin.H{
		"name":  req.Name,
		"roles": req.Roles,
	})
}

func ValidateEmployeePostRequest(req EmployeePostRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if len(req.Roles) == 0 {
		return errors.New("roles is required")
	}

	if !models.AreAllRolesValid(req.Roles) {
		return errors.New("invalid role")
	}

	return nil
}
