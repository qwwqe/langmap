package langmap

import "github.com/go-gorp/gorp"

type BaseService struct {
	Engine *Engine
	Prefix string
}

func (s *BaseService) Db() *gorp.DbMap {
	return s.Engine.DB
}
