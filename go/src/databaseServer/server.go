package server

import "simpleDB/database"

// Write Read describe the actions of the protocol
const (
	Write = iota + 1
	Read
)

// RequestPayload describes the data sent by each request
type RequestPayload struct {
	Action int      `json:"action"`
	args   []string `json:"args"`
}

// Server is the struct which interfaces with the network
type Server struct {
	db database.SimpleDB
}

// NewServer is the server constructor
func NewServer(db database.SimpleDB) *Server {

	return &Server{
		db: db,
	}
}

func (s *Server) Protocol(requestPayload RequestPayload) (string, error) {

	var err error
	var val string
	switch requestPayload.Action {
	case Write:
		err = s.db.Write(requestPayload.args[0], requestPayload.args[1])
		if err != nil {
			return "", err
		}
	case Read:
		val, err = s.db.FastGet(requestPayload.args[0])
		if err != nil {

			return "", err
		}

	}
	return val, nil
}
