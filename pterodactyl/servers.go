package pterodactyl

import "errors"

var (
	ServerNotFoundErr = errors.New("server not found")
)

func GetServer(serverID string) (*Server, error) {
	var (
		server *Server
	)
	for _, s := range Servers {
		if s.Config.ServerID == serverID {
			server = s
			return server, nil
		}
	}
	return nil, ServerNotFoundErr
}
