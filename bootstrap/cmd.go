package bootstrap

func Run() {
	s := NewServer(":3000")
	s.SetUpRoutes()
	s.RunServer()
}
