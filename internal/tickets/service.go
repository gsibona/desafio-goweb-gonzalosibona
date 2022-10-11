package tickets

import (
	"desafio-goweb-gonzalosibona/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface{
	GetTotalTickets(*gin.Context,string)([]domain.Ticket,error)
	AverageDestination(*gin.Context,string)(float64,error)
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service{
	return &service{
		repository: r,
	}
}

func (s *service)GetTotalTickets(c *gin.Context,dest string) ([]domain.Ticket,error) {
	return s.repository.GetTicketByDestination(c,dest)
}

func (s *service)AverageDestination(c *gin.Context,dest string) (float64,error) {
	tickets,err:= s.repository.GetAll(c)
	if err != nil {
		return 0.0, err
	}
	ticketsDest,err:= s.repository.GetTicketByDestination(c,dest)
	if err != nil {
		return 0.0, err
	}
	return float64(len(ticketsDest))/float64(len(tickets))*100,nil
}