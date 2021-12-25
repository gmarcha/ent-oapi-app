package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmarcha/goswagger-ent-app/ent"
	"github.com/gmarcha/goswagger-ent-app/ent/event"
	"github.com/gmarcha/goswagger-ent-app/ent/user"
	"github.com/gmarcha/goswagger-ent-app/internal/api/errors"
	"github.com/google/uuid"
)

func (a Server) ListEvent(c *gin.Context) {

	res, err := a.Client.Event.Query().
		Order(ent.Asc(event.FieldDateStart)).
		All(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.ErrJSON(http.StatusInternalServerError, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) CreateEvent(c *gin.Context) {

	var ev ent.Event

	if err := c.ShouldBind(&ev); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := a.Client.Event.Create().
		SetName(ev.Name).
		SetDateStart(ev.DateStart).
		SetDateEnd(ev.DateEnd).
		SetNbTutors(ev.NbTutors).
		Save(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.ErrJSON(http.StatusInternalServerError, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) DeleteEvent(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	if err := a.Client.Event.DeleteOneID(uuid).Exec(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

func (a Server) ReadEvent(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := a.Client.Event.Query().
		Where(event.ID(uuid)).
		Only(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) UpdateEvent(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	ev, err := a.Client.Event.Query().
		Where(event.ID(uuid)).
		Only(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	if err := c.ShouldBind(ev); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := ev.Update().
		SetName(ev.Name).
		SetDateStart(ev.DateStart).
		SetDateEnd(ev.DateEnd).
		SetNbTutors(ev.NbTutors).
		Save(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.ErrJSON(http.StatusInternalServerError, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) UnsubscribeEvent(c *gin.Context, id string, userId string) {

	ev_uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	user_uuid, err := uuid.Parse(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	ev, err := a.Client.Event.Query().
		Where(
			event.ID(ev_uuid),
			event.HasUsersWith(user.ID(user_uuid)),
		).
		Only(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	res, err := ev.Update().
		RemoveUserIDs(user_uuid).
		Save(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, errors.ErrJSON(http.StatusInternalServerError, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) SubscribeEvent(c *gin.Context, id string, userId string) {

	ev_uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	user_uuid, err := uuid.Parse(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	ev, err := a.Client.Event.Query().
		Where(event.ID(ev_uuid)).
		Only(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	res, err := ev.Update().
		AddUserIDs(user_uuid).
		Save(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (a Server) ListEventUsers(c *gin.Context, id string) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrJSON(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := a.Client.Event.Query().
		Where(event.ID(uuid)).
		QueryUsers().
		All(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, errors.ErrJSON(http.StatusNotFound, err.Error()))
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
