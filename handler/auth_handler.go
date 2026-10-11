package handler

import (
	"net/http"
	"strings"

	"github.com/vn-contrib/urhome-server/contexts"
	"github.com/vn-contrib/urhome-server/model"
	"github.com/vn-contrib/urhome-server/pkg/bar"
)

type AuthHandler struct {
	Next http.Handler
}

var _ http.Handler = (*AuthHandler)(nil)

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authCtx := contexts.WithAuth(ctx, func() *model.User {
		idToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		token, err := contexts.UseFirebaseAuth(ctx).VerifyIDToken(ctx, idToken)
		if err != nil {
			panic(err)
		}

		u := &model.User{FirebaseUID: token.UID}
		if err := bar.FindOrCreateBy(ctx, u, "firebase_uid"); err != nil {
			panic(err)
		}

		return u
	})

	h.Next.ServeHTTP(w, r.WithContext(authCtx))
}
