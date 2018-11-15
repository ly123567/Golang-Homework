package handlers

import "net/http"

// StaticHandler return static file content.
func StaticHandler(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}

// SubmitHandler return the form content.
func SubmitHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	text := `
		<!DOCTYPE html>
		<html>
		<head>
    		<meta charset="utf-8">
    		<title>Submit Result</title>
  		</head>
		<body>
		` + "<div>Hello! " + req.Form.Get("name") + "</div>" +
		"<div>Your email: " + req.Form.Get("email") + "</div>" +
		`</body>
		</html>`
	w.Write([]byte(text))
}

//UnknownHandler return a 500 HTTP code.
func UnknownHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("500 Internal Error"))
}
