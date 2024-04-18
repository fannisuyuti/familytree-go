package serializer

import (
	"github.com/gin-gonic/gin"
	"suyuti.com/famtrees/model"
)

type PersonSerializer struct {
	C *gin.Context
}

type PersonResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Father uint   `json:"father"`
	Mother uint   `json:"mother"`
	Birth  string `json:"birth"`
	Death  string `json:"death"`
	Gender string `json:"gender"`
}

func (person *PersonSerializer) Response() PersonResponse {
	p := person.C.MustGet("person").(model.Person)
	return PersonResponse{
		ID:     p.ID,
		Name:   p.Name,
		Father: p.Father,
		Mother: p.Mother,
		Birth:  p.Birth,
		Death:  p.Death,
		Gender: p.Gender,
	}
}
