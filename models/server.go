package models

type Server struct {
	clients map[*Client]string
}

func (s *Server) GetClient(id string) *Client {
	for cli, cliID := range s.clients {
		if cliID == id {
			return cli
		}
	}

	return nil
}

func (s *Server) AddClient(id string) *Client {
	if s.clients == nil {
		s.clients = make(map[*Client]string)
	}

	if s.GetClient(id) != nil {
		return nil
	}

	var cli *Client = new(Client)
	cli.id = id
	return cli
}

func (s *Server) VerifyPassword(pw string, id string) bool {
	for cli, cliID := range s.clients {
		if cliID == id {
			if cli.password == pw {
				return true
			}
		}
	}

	return false
}
