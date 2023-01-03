package cmd

import (
	s "github.com/deall-users/pkg/server"
)

// init server and its services
func App() (*s.Server, error) {
	server, err := s.NewServer()
	if err != nil {
		return nil, err
	}

	server.InitializeServices()
	return server, nil
}
