package api

import (
	"guideventureapi/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	db         db.Database
}

func NewServer(listenAddr string) (*Server, error) {
	db, err := db.NewSQLiteDb()
	if err != nil {
		return nil, err
	}

	// err = sqliteDb.CreateDummyData()
	// if err != nil {
	// 	log.Panic(err)
	// }

	return &Server{
		listenAddr: listenAddr,
		db:         db,
	}, nil
}

func (s *Server) GetGames(c *gin.Context) {
	games, err := s.db.GetGames()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, &games)

}

func (s *Server) GetGame(c *gin.Context) {
	id := c.Param("id")
	game, err := s.db.GetGame(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &game)
}

func (s *Server) Start() error {
	router := gin.Default()

	router.GET("/games/", s.GetGames)
	router.GET("/games/:id/", s.GetGame)

	return router.Run(s.listenAddr)
}
