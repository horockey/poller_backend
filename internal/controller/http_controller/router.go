package http_controller

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (ctrl *httpController) newRouter() *mux.Router {
	routes := Routes{
		Route{
			"AttemptPollIdAttemptIdDelete",
			strings.ToUpper("Delete"),
			"/attempt/{poll_id}/{attempt_id}",
			ctrl.AttemptPollIdAttemptIdDelete,
		},

		Route{
			"AttemptPollIdAttemptIdGet",
			strings.ToUpper("Get"),
			"/attempt/{poll_id}/{attempt_id}",
			ctrl.AttemptPollIdAttemptIdGet,
		},

		Route{
			"AttemptPollIdDelete",
			strings.ToUpper("Delete"),
			"/attempt/{poll_id}",
			ctrl.AttemptPollIdDelete,
		},

		Route{
			"AttemptPollIdGet",
			strings.ToUpper("Get"),
			"/attempt/{poll_id}",
			ctrl.AttemptPollIdGet,
		},

		Route{
			"PruneAttemptsDelete",
			strings.ToUpper("Delete"),
			"/prune_attempts",
			ctrl.PruneAttemptsDelete,
		},

		Route{
			"PollGet",
			strings.ToUpper("Get"),
			"/poll",
			ctrl.PollGet,
		},

		Route{
			"PollIdDelete",
			strings.ToUpper("Delete"),
			"/poll/{id}",
			ctrl.PollIdDelete,
		},

		Route{
			"PollIdGet",
			strings.ToUpper("Get"),
			"/poll/{id}",
			ctrl.PollIdGet,
		},

		Route{
			"PollPost",
			strings.ToUpper("Post"),
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

	return router
}
