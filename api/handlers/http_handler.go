package handlers

import (
	"reminder_bot/api/services"
)

type HTTPHandler struct {
	Service *services.MyService
}
