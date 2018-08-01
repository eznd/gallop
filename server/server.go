package server

import (
	"bytes"
	"eznd/gallop/log"
	"eznd/gallop/util"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// Server struct for storing current server information
type Server struct {
	Router   *gin.Engine
	RootPath string
	Host     string
	Port     string
}

// New function create server wrapper
func New(rootPath string, host string, port string) *Server {
	return &Server{
		Router:   gin.Default(),
		RootPath: rootPath,
		Host:     host,
		Port:     port,
	}
}

// StartRouter starts REST server
func (s *Server) StartRouter() {
	s.Router.GET("/reports/", s.handleStatic)
	s.Router.Use(static.Serve("/reports", static.LocalFile(s.RootPath+"static/reports/", true)))
	s.Router.POST("/create", s.createReport)

	err := s.Router.Run(s.Host + ":" + s.Port)
	fmt.Print(err)
}

func (s *Server) handleStatic(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func (s *Server) createReport(c *gin.Context) {
	data, _ := c.GetRawData()
	if string(data) == "" {
		c.JSON(http.StatusNoContent, "")
	}

	dirName := util.RandStringBytes(12)
	sourceDir := "static/source/" + dirName
	reportDir := "static/reports/" + dirName

	err := os.MkdirAll(s.RootPath+sourceDir, os.ModeDir|os.ModePerm)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "")
	}

	err = os.MkdirAll(s.RootPath+reportDir, os.ModeDir|os.ModePerm)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "")
	}

	err = ioutil.WriteFile(s.RootPath+sourceDir+"/report.xml", data, 0644)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "")
	}

	cmd := exec.Command("allure", "generate", sourceDir, "-o", reportDir)
	cmd.Dir = s.RootPath
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "")
	}

	// docker patch
	host := s.Host
	if host == "" {
		host = "localhost"
	}

	c.Header("Location", fmt.Sprintf("http://%s:%s/reports/%s", host, s.Port, dirName))
	c.Status(http.StatusCreated)
}
