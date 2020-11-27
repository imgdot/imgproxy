package main

import "net/http"

var landingTmpl = []byte(`
<!doctype html>
<html>
	<head><title>This serivce is provided by ImgDot</title></head>
	<body>
		<h1>This serivce is provided by ImgDot</h1>
		<p style="font-size:1.2em">Please visit: <a href="https://imgdot.dev" target="_blank">https://imgdot.dev</a> for details</p>
	</body>
</html>
`)

func handleLanding(reqID string, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Tyle", "text/html")
	rw.WriteHeader(200)
	rw.Write(landingTmpl)
}
