package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"grader/configs"
	"grader/pkg/exercise"
	"net/http"
)

type Rest struct {
	addr string
	*chi.Mux

	app IApp
}

type IApp interface {
	GetExercise(id string) (exercise.Exercise, error)
	CheckExercise(e exercise.Exercise) error
}

func NewRest(app IApp, config configs.ServerConfig) *Rest {
	rest := &Rest{
		addr: config.Host,
		app:  app,
	}

	r := chi.NewRouter()
	r.Route("/v1/api", func(rapi chi.Router) {
		r.Post("/exercise/{ID}/upload", rest.handleUploadExercise) // upload file with exercise
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Get("/exercise", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("exercises list"))
	})
	r.Get("/exercise/{ID}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("exercise"))
	})

	rest.Mux = r

	return rest
}

func (rest *Rest) Serve() {
	log.Debug().Msg("Starting REST exercise at [" + rest.addr + "]")

	if err := http.ListenAndServe(rest.addr, rest); err != nil {
		log.Fatal().Err(err).Msg("Failed to start REST exercise")
	}
}

func (rest *Rest) handleUploadExercise(w http.ResponseWriter, r *http.Request) {
	//eID := chi.URLParam(r, "ID")
	//user
	//
	//user from ctx
	//exerciseID from url param
	//listFiles
	//
	//
	//rest.app.CheckExercise()
	//rest.app.
	//	log.Debug().Msg("Uploading exercise")
}
