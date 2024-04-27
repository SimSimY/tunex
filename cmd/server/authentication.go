package main

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

type AppUser struct {
	OrgName string `json:"org_name"`
}

func (u *AppUser) String() string {
	return fmt.Sprintf("%s", u.OrgName)
}

// Middleware function, which will be called for each request
func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "user", &AppUser{OrgName: "simon"})
		next.ServeHTTP(w, r)

	})
}
