package patients

import (
	"net/http"

	"bitbucket.org/topdoctors/tools/controllers"
	"github.com/gin-gonic/gin"

	dto "basic-go-training/internal/domain/dtos/patients"
	service "basic-go-training/internal/domain/services/patients"
)

type Server struct {
	PatientService service.PatientsService
}

func NewServer(router *gin.Engine) {
	server := &Server{
		PatientService: service.NewService(),
	}
	server.registerEndpoints(router)
}

func (s *Server) registerEndpoints(router *gin.Engine) {
	router.GET("/patients", s.Search)
	router.GET("/patients/:id", s.Get)
	router.POST("/patients", s.Create)
	router.PUT("/patients/:id", s.Update)
	router.DELETE("/patients/:id", s.Delete)
}

func (s *Server) Search(ctx *gin.Context) {
	req := new(dto.SearchRequest)
	if err := controllers.ParseRequest(ctx, req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reply, err := s.PatientService.Search(req)
	if err != nil {
		controllers.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, reply)
}

func (s *Server) Get(ctx *gin.Context) {
	req := new(dto.GetRequest)
	if err := controllers.ParseRequest(ctx, req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reply, err := s.PatientService.Get(req)
	if err != nil {
		controllers.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, reply)
}

func (s *Server) Create(ctx *gin.Context) {
	req := new(dto.CreateRequest)
	if err := controllers.ParseRequest(ctx, req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reply, err := s.PatientService.Create(req)
	if err != nil {
		controllers.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, reply)
}

func (s *Server) Update(ctx *gin.Context) {
	req := new(dto.UpdateRequest)
	if err := controllers.ParseRequest(ctx, req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reply, err := s.PatientService.Update(req)
	if err != nil {
		controllers.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, reply)
}

func (s *Server) Delete(ctx *gin.Context) {
	req := new(dto.DeleteRequest)
	if err := controllers.ParseRequest(ctx, req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reply, err := s.PatientService.Delete(req)
	if err != nil {
		controllers.Error(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, reply)
}
