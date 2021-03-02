package handle

import (
	"net/http"
)

func HTTPInterceptor(h http.HandlerFunc)http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")
		token := r.Form.Get("token")
		//验证token的有效性
		if len(username) < 3 || len(token) < 40 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	})
}
