package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"grader/configs"
	"grader/pkg/exercise"
	"grader/pkg/user"
	"grader/web"
	"html/template"
	"io"
	"net/http"
)

type Rest struct {
	addr string
	*chi.Mux
	tmpl *template.Template

	app IApp
}

type IApp interface {
	GetExercise(user user.User, id string) (exercise.Exercise, error)
	CheckExercise(user user.User, e exercise.Exercise) error
}

func NewRest(app IApp, config configs.ServerConfig) *Rest {
	tmpl, err := web.ParseTemplates()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse src")
	}

	rest := &Rest{
		addr: config.Host,
		app:  app,
		tmpl: tmpl,
	}

	r := chi.NewRouter()
	r.Route("/v1/api", func(rapi chi.Router) {})

	r.Get("/", rest.index)
	r.Get("/exercise", rest.exercisesList)
	r.Get("/exercise/{ID}", rest.exerciseView)
	r.Post("/exercise/{ID}/upload", rest.handleUploadExercise)

	r.Handle("/assets/*", web.StaticFiles(""))
	rest.Mux = r

	return rest
}

func (rest *Rest) Serve() {
	log.Debug().Msg("Starting REST exercise at [" + rest.addr + "]")

	if err := http.ListenAndServe(rest.addr, rest); err != nil {
		log.Fatal().Err(err).Msg("Failed to start REST exercise")
	}
}

func (rest *Rest) index(w http.ResponseWriter, r *http.Request) {
	rest.tmpl.ExecuteTemplate(w, "index.html", nil)
}
func (rest *Rest) exercisesList(w http.ResponseWriter, r *http.Request) {
	rest.tmpl.ExecuteTemplate(w, "index.html", nil)
}
func (rest *Rest) exerciseView(w http.ResponseWriter, r *http.Request) {
	usr := user.FromContext(r.Context())
	if usr == nil {
		noAuth(w, r)

		return
	}

	rest.tmpl.ExecuteTemplate(w, "exercise.html", nil)
}

func (rest *Rest) handleUploadExercise(w http.ResponseWriter, r *http.Request) {
	usr := user.FromContext(r.Context())
	if usr == nil {
		noAuth(w, r)

		return
	}

	rest.tmpl.ExecuteTemplate(w, "exercise.html", nil)
}

func noAuth(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("no auth")

	w.WriteHeader(http.StatusUnauthorized)
	io.WriteString(w, "вы не авторизованы")
}
