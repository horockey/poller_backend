package http_controller

import (
	"embed"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed docs/*
var docs embed.FS

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (ctrl *httpController) newRouter() *mux.Router {
	routes := Routes{
		{
			"AttemptPollIdAttemptIdDelete",
			http.MethodDelete,
			"/attempt/{poll_id}/{attempt_id}",
			ctrl.AttemptPollIdAttemptIdDelete,
		},
		{
			"AttemptPollIdAttemptIdGet",
			http.MethodGet,
			"/attempt/{poll_id}/{attempt_id}",
			ctrl.AttemptPollIdAttemptIdGet,
		},
		{
			"AttemptPollIdDelete",
			http.MethodDelete,
			"/attempt/{poll_id}",
			ctrl.AttemptPollIdDelete,
		},
		{
			"AttemptPollIdGet",
			http.MethodGet,
			"/attempt/{poll_id}",
			ctrl.AttemptPollIdGet,
		},
		{
			"PruneAttemptsDelete",
			http.MethodDelete,
			"/prune_attempts",
			ctrl.PruneAttemptsDelete,
		},
		{
			"PollGet",
			http.MethodGet,
			"/poll",
			ctrl.PollGet,
		},
		{
			"PollIdDelete",
			http.MethodDelete,
			"/poll/{poll_id}",
			ctrl.PollIdDelete,
		},
		{
			"PollIdGet",
			http.MethodGet,
			"/poll/{poll_id}",
			ctrl.PollIdGet,
		},
		{
			"PollPost",
			http.MethodPost,
			"/poll",
			ctrl.PollPost,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.
		Methods(http.MethodGet).
		Path("/").
		Name("Docs").
		Handler(http.FileServer(http.FS(docs)))

	return router
}
