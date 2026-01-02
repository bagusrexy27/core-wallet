package handler

type CheckHealthHandler interface {
	CheckHealth() error
}

type checkHealthHandler struct{}

func NewCheckHealthHandler() CheckHealthHandler {
	return &checkHealthHandler{}
}

func (h *checkHealthHandler) CheckHealth() error {
	return nil
}
