package internal

import "net/http"

func ApiDocsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!doctype html>
	<html>
		<head>
			<title>API Reference</title>
			<meta charset="utf-8" />
			<meta
			  name="viewport"
			  content="width=device-width, initial-scale=1" />
		</head>
		<body>
			<script
			  id="api-reference"
			  data-url="/openapi.json"></script>
			<script src="http://cdn.local/scalar.js"></script>
		</body>
	</html>`))
}
