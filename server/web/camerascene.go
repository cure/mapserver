package web

import (
	"mapserver/app"
	"net/http"
)

type CameraScene struct {
	ctx *app.App
}

func (t *CameraScene) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//str := strings.TrimPrefix(req.URL.Path, "/api/tile/")

	/*
		parts := strings.Split(str, "/")
		if len(parts) != 4 {
			resp.WriteHeader(500)
			resp.Write([]byte("wrong number of arguments"))
			return
		}
	*/

	//x, _ := strconv.Atoi(parts[0])
	//y, _ := strconv.Atoi(parts[1])
	//z, _ := strconv.Atoi(parts[2])

	var err error

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))

	} else {
		resp.Header().Add("content-type", "image/png")
		resp.Write([]byte{})
	}
}
