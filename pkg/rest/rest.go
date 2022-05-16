package rest

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"grader/configs"
	"grader/pkg/exercise"
	"grader/pkg/user"
	"grader/web"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Rest struct {
	addr string
	*chi.Mux
	tmpl     *template.Template
	filesDir string

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
		addr:     config.Host,
		app:      app,
		tmpl:     tmpl,
		filesDir: config.FilesDir,
	}

	r := chi.NewRouter()
	logMdlw := func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info().Str("method", r.Method).Str("url", r.URL.String()).Msg("request")
			handler.ServeHTTP(w, r)
		})
	}

	r.Route("/v1/api", func(rapi chi.Router) {})
	r.Group(func(r chi.Router) {
		r.Use(logMdlw)

		r.Get("/", rest.index)
		r.Get("/exercise", rest.exercisesList)
		r.Get("/exercise/{ID}", rest.exerciseView)
		r.Post("/exercise/{ID}/upload", rest.handleUploadExercise)
	})
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

	if err := rest.tmpl.ExecuteTemplate(w, "exercise.html", usr); err != nil {
		log.Err(err).Msg("failed to execute template")
	}
}

const maxFilesSize = 10 * 1024 * 1024 // 10MB
var ErrNoFiles = errors.New("no files")

func (rest *Rest) handleUploadExercise(w http.ResponseWriter, r *http.Request) {
	exID := chi.URLParam(r, "ID")
	if exID == "" {
		log.Err(ErrNoFiles).Msg("no exercise id")
		http.Error(w, ErrNoFiles.Error(), http.StatusBadRequest)

		return
	}

	usr := user.FromContext(r.Context())
	if usr == nil {
		noAuth(w, r)

		return
	}

	if err := r.ParseMultipartForm(maxFilesSize); err != nil {
		log.Err(err).Msg("failed to parse multipart form")

		return
	}
	defer r.MultipartForm.RemoveAll()

	formFiles, ok := r.MultipartForm.File["exerciseFiles"]
	if !ok {
		log.Err(ErrNoFiles).Msg("failed to get exercise files")

		return
	}

	numTry := "1" // номер попытки

	files := make([]exercise.File, len(formFiles))

	for i, v := range formFiles {
		p, err := filepath.Abs(filepath.Join(rest.filesDir, exID, usr.ID, numTry))
		if err != nil {
			log.Err(err).Msg("failed to get absolute path")

			return
		}

		file, err := rest.copyFile(v, p)
		if err != nil {
			log.Err(err).Msg("failed to copy file")

			return
		}

		files[i] = file
	}

	if err := rest.app.CheckExercise(*usr, exercise.Exercise{
		ID:     exID,
		UserID: usr.ID,
		Files:  files,
	}); err != nil {
		log.Err(err).Msg("failed to check exercise")

		return
	}

	rest.tmpl.ExecuteTemplate(w, "exercise.html", nil)
}

var ErrWrongExtension = errors.New("wrong extension")

func (rest *Rest) copyFile(fh *multipart.FileHeader, to string) (exercise.File, error) {
	if !strings.HasSuffix(fh.Filename, ".go") {
		return exercise.File{}, fmt.Errorf("file [%s] error: %w", fh.Filename, ErrWrongExtension)
	}

	f, err := fh.Open()
	if err != nil {
		return exercise.File{}, err
	}
	defer f.Close()

	if err = os.MkdirAll(to, os.ModePerm); err != nil {
		return exercise.File{}, fmt.Errorf("failed to create file: %w", err)
	}

	fpath := filepath.Base(fh.Filename) // or generate unique name

	nf, err := os.Create(to + "/" + fpath)
	if err != nil {
		return exercise.File{}, fmt.Errorf("failed to create file: %w", err)
	}

	if _, err = io.Copy(nf, f); err != nil {
		return exercise.File{}, fmt.Errorf("failed to copy file: %w", err)
	}

	ef := exercise.File{
		Name: fpath,
		Path: nf.Name(),
	}

	return ef, nil
}

func noAuth(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("no auth")

	w.WriteHeader(http.StatusUnauthorized)
	io.WriteString(w, "вы не авторизованы")
}
