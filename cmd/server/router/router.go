package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/ticket"
	"github.com/gin-gonic/gin"
)

type Router struct {
	srv  *gin.Engine
	list []domain.Ticket
}

func NewRouter(s *gin.Engine, l []domain.Ticket) *Router {
	return &Router{
		srv:  s,
		list: l,
	}
}

func (rou *Router) MapRoutes() {
	repo := tickets.NewRepository(rou.list)
	service := tickets.NewService(repo)
	handler := handler.NewService(service)
	ticketRoutes := rou.srv.Group("/ticket")
	{
		ticketRoutes.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
		ticketRoutes.GET("/getAverage/:dest", handler.AverageDestination())
	}
}
