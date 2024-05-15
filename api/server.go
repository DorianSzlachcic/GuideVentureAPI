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
	gameId := c.Param("gameId")
	game, err := s.db.GetGame(gameId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &game)
}

func (s *Server) GetSteps(c *gin.Context) {
	gameId := c.Param("gameId")
	steps, err := s.db.GetSteps(gameId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &steps)
}

func (s *Server) GetStep(c *gin.Context) {
	gameId := c.Param("gameId")
	stepIndex := c.Param("stepIndex")
	step, err := s.db.GetStep(gameId, stepIndex)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &step)
}

func (s *Server) GetQuestions(c *gin.Context) {
	gameId := c.Param("gameId")
	stepIndex := c.Param("stepIndex")
	questions, err := s.db.GetQuestions(gameId, stepIndex)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &questions)
}

func (s *Server) Start() error {
	router := gin.Default()

	router.GET("/games/", s.GetGames)
	router.GET("/games/:gameId/", s.GetGame)
	router.GET("/games/:gameId/steps/", s.GetSteps)
	router.GET("/games/:gameId/steps/:stepIndex", s.GetStep)
	router.GET("/games/:gameId/steps/:stepIndex/questions", s.GetQuestions)

	return router.Run(s.listenAddr)
}
