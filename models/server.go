package models

type Server struct {
	Clients map[*Client]string
}

func (s *Server) GetClient(id string) *Client {
	for cli := range s.Clients {
		if cli.GetID() == id {
			return cli
		}
	}

	return nil
}

func (s *Server) AddClient(id string) *Client {
	if s.Clients == nil {
		s.Clients = make(map[*Client]string)
	}

	if s.GetClient(id) != nil {
		return nil
	}

	var cli = new(Client)
	cli.id = id
	return cli
}

func (s *Server) VerifyPassword(pw string, id string) bool {
	for cli, cliPassword := range s.Clients {
		if cli.GetID() == id {
			if cliPassword == pw {
				return true
			}
		}
	}

	return false
}
