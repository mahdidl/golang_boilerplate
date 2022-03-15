package Ticket

import "github.com/google/uuid"

type CreateTicketRequest struct {
	UserId   uuid.UUID `json:"userId"`
	UserName string    `json:"userName" validate:"required,min=3"`
	Subject  string    `json:"subject" validate:"required"`
	Message  string    `json:"message" validate:"required"`
	Like     bool      `json:"like" validate:"required"`
	Image    []byte    `json:"image"`
}