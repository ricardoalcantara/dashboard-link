package foo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/dashboard-link/internal/domain"
	"github.com/ricardoalcantara/dashboard-link/internal/models"
	"github.com/ricardoalcantara/dashboard-link/internal/utils"
)

func post(c *gin.Context) {
	var input FooRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		out := utils.GetValidationErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Form Validation", Details: out})
		return
	}

	foo := models.Foo{
		ID:   models.NewDbUUID(),
		Name: input.Name,
	}

	err := foo.Save()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": utils.PrintError(err)})
		return
	}

	c.JSON(http.StatusCreated, FooView{
		Id:   foo.ID.String(),
		Name: foo.Name,
	})
}
