package server

type Server struct {
	Address string
}

func New() *Server {
	return &Server{}
}

func Default() *Server {
	return &Server{
		Address: "127.0.0.1:2004",
	}
}

func (s *Server) Run() error {
	return nil
}
