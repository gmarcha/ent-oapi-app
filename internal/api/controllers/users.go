package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmarcha/goswagger-ent-app/ent"
	"github.com/gmarcha/goswagger-ent-app/ent/user"
	"github.com/gmarcha/goswagger-ent-app/internal/api/errors"
	"github.com/google/uuid"
)

func (a Server) ListUser(c *gin.Context) {

	res, err := a.Client.User.Query().
		Order(ent.Asc(user.FieldName)).
		All(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.ErrJSON(http.StatusInternalServerError, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) CreateUser(c *gin.Context) {

	var u ent.User

	if err := c.ShouldBind(&u); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := a.Client.User.Create().
		SetName(u.Name).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		Save(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.ErrJSON(http.StatusInternalServerError, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) DeleteUser(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	if err := a.Client.User.DeleteOneID(uuid).Exec(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

func (a Server) ReadUser(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := a.Client.User.Query().
		Where(user.ID(uuid)).
		Only(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) UpdateUser(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	u, err := a.Client.User.Query().
		Where(user.ID(uuid)).
		Only(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	if err := c.ShouldBind(u); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := u.Update().
		SetName(u.Name).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetTutor(u.Tutor).
		Save(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.ErrJSON(http.StatusInternalServerError, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) ListUserEvents(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := a.Client.User.Query().
		Where(user.ID(uuid)).
		QueryEvents().
		All(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
