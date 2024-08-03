package schemas

type ServerIsAlive struct {
	Status string
}

func NewServerIsAlive() *ServerIsAlive {
	return &ServerIsAlive{
		Status: "ok",
	}
}
