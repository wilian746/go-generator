package server

type Interface interface {
	CreateFoldersAndFiles(pathDestiny, moduleName string) error
}

type Server struct{}

func NewServer() Interface {
	return &Server{}
}

func (s *Server) CreateFoldersAndFiles(pathDestiny, moduleName string) error {
	return nil
}
