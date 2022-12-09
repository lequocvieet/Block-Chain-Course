package middlewares

import (
	"errors"
	jwt "golang_service/auth"
	res "golang_service/utils"
	"net/http"
)

func CheckJwt(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := jwt.Verify(r)

		if err != nil {
			res.ERROR(w, 401, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}
