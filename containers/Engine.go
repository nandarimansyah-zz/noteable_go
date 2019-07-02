package containers

import (
	"os"

	"github.com/nandarimansyah/noteable_go/controllers"
	"github.com/nandarimansyah/noteable_go/infrastructures"
	"github.com/nandarimansyah/noteable_go/repositories"
	"github.com/nandarimansyah/noteable_go/routes"
	"github.com/nandarimansyah/noteable_go/services"
)

// Engine is the struct of container service
type Engine struct {
	MongoDB *infrastructures.MongoDB
}

// NewEngine create a new container of the app
func NewEngine() *Engine {
	e := &Engine{}

	mongoDb := infrastructures.NewMongoDB(os.Getenv("mongo_url"), os.Getenv("mongo_database"))
	e.MongoDB = mongoDb

	return e
}

// Make handle application initalization
func (e *Engine) Make(route *routes.Route) *routes.Route {

	notesRepository := new(repositories.NotesRepository)
	notesRepository.MongoDB = e.MongoDB.DBCached

	notesService := new(services.NotesService)
	notesService.NotesRepository = notesRepository

	notesController := new(controllers.NotesController)
	notesController.NotesService = notesService

	commentRepo := new(repositories.CommentRepository)
	commentRepo.MongoDB = e.MongoDB.DBCached

	commentService := new(services.CommentService)
	commentService.Repository = commentRepo

	commentController := new(controllers.CommentController)
	commentController.Service = commentService

	memberRepo := new(repositories.MemberRepository)
	memberService := new(services.MemberService)
	memberService.Repository = memberRepo

	memberController := new(controllers.MemberController)
	memberController.Service = memberService

	route.NotesController = notesController
	route.CommentController = commentController
	route.MemberController = memberController

	return route
}
