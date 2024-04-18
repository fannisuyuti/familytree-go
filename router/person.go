package router

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"suyuti.com/famtrees/common"
	"suyuti.com/famtrees/model"
	"suyuti.com/famtrees/serializer"
)

func Tree(router *gin.RouterGroup) {
	router.GET("/", GetTree)
	router.GET("/add", AddForm)
	router.POST("/add", AddPerson)
	router.GET("/:id", UpdateForm)
}

func GetTree(c *gin.Context) {
	persons, err := model.GetAll()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	var personResponses []model.PersonResponse

	for _, person := range persons {
		image := "/static/primitives/icon/male.png"
		if person.Gender == "F" {
			image = "/static/primitives/icon/female.png"
		}

		description := "Lahir :" + person.Birth + "Meninggal : " + person.Death
		if person.Death == "" {
			description = "Lahir :" + person.Birth
		}

		personResponses = append(personResponses, model.PersonResponse{
			Id:          person.ID,
			Parents:     []uint{person.Father, person.Mother},
			Title:       person.Name,
			Label:       person.Name,
			Description: description,
			Image:       image,
		})
	}

	sort.Slice(personResponses, func(i, j int) bool {
		return personResponses[i].Id < personResponses[j].Id
	})

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"pohon": personResponses,
	})
}

func AddForm(c *gin.Context) {
	persons, err := model.GetAll()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	ayahs := []model.Option{}
	ibus := []model.Option{}

	for _, person := range persons {
		if person.Gender == "M" {
			ayahs = append(ayahs, model.Option{
				ID:   person.ID,
				Name: person.Name,
			})
		} else {
			ibus = append(ibus, model.Option{
				ID:   person.ID,
				Name: person.Name,
			})
		}
	}

	c.HTML(http.StatusOK, "form.tmpl", gin.H{
		"ayah": ayahs,
		"ibu":  ibus,
	})
}

func UpdateForm(c *gin.Context) {
	var selectedPerson model.Person

	persons, err := model.GetAll()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	ayahs := []model.Option{}
	ibus := []model.Option{}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("param", err))
		return
	}

	for _, person := range persons {

		if person.ID == uint(id) {
			selectedPerson = person
		}

		if person.Gender == "M" {
			ayahs = append(ayahs, model.Option{
				ID:   person.ID,
				Name: person.Name,
			})
		} else {
			ibus = append(ibus, model.Option{
				ID:   person.ID,
				Name: person.Name,
			})
		}
	}

	c.HTML(http.StatusOK, "update.tmpl", gin.H{
		"ayah":         ayahs,
		"ibu":          ibus,
		"personId":     selectedPerson.ID,
		"personName":   selectedPerson.Name,
		"personFather": selectedPerson.Father,
		"personMother": selectedPerson.Mother,
		"personBirth":  selectedPerson.Birth,
		"personDeath":  selectedPerson.Death,
		"personGender": selectedPerson.Gender,
	})
}

func AddPerson(c *gin.Context) {
	var person model.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("person", err))
		return

	}

	if person.ID > 0 {
		UpdatePerson(c)
		return
	}

	if err := model.SaveOne(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("person", person)
	serializer := serializer.PersonSerializer{C: c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UpdatePerson(c *gin.Context) {
	var person model.Person
	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("person", err))
		return

	}

	oldPerson, err := model.GetOne(person.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	if err := oldPerson.Update(&person); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("person", person)
	serializer := serializer.PersonSerializer{C: c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
