package auth

import "net/http"

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /login/{provider}", Login)
	mux.HandleFunc("GET /callback/{provider}", LoginCallback)
	mux.HandleFunc("GET /logout/{provider}", Logout)
	mux.HandleFunc("GET /me", GetUserSession)

	return mux
}
