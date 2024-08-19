package api

import (
	"guideventureapi/api/middleware"
	"guideventureapi/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	db         db.Database
}

type Option func(*Server) error

func NewServer(options ...Option) (*Server, error) {
	server := &Server{}
	for _, opt := range options {
		if err := opt(server); err != nil {
			return nil, err
		}
	}
	return server, nil
}

func WithListenAddr(listenAddr string) Option {
	return func(s *Server) error {
		s.listenAddr = listenAddr
		return nil
	}
}

func WithDatabase(db db.Database, err error) Option{
	return func(s *Server) error {
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

func WithDummyData() Option {
	return func(s *Server) error {
		return s.db.CreateDummyData()
	}
}

func (s *Server) Start() error {
	router := gin.Default()
	router.Use(middleware.ErrorHandler)

	router.GET("/games/", s.GetGames)
	router.GET("/games/:gameId/", s.GetGame)
	router.GET("/games/:gameId/steps/", s.GetSteps)
	router.GET("/games/:gameId/steps/:stepIndex", s.GetStep)

	return router.Run(s.listenAddr)
}
