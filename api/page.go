package api

import (
	"html/template"
	"log"
	"net/http"
)

// Page renders the homepage with the IP address of the request within the body
// so the page doesn't have to call the API after it loads.
func Page(w http.ResponseWriter, r *http.Request) {
	reply, err := identifyIP(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := page.Execute(w, reply); err != nil {
		log.Printf("rendering page: %v", err)
	}
}

var page = template.Must(template.New("page").Parse(`<!DOCTYPE html>

<html lang="en">

<head>
    <title>whatsmyip</title>
    <meta charset="utf-8">
    <meta name="theme-color" content="#fff" />
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="Golang API running as a Vercel Serverless Function which returns your public IP address">
</head>

<body>
    <div style="text-align: center;">
        <div>
            <h1><div><a href="https://github.com/wafer-bw/whatsmyip">wafer-bw/whatsmyip</a></div></h1>
            <h2>Your public IP address is: <div id="ip">{{.Ip}}</div></h2>
            <div>
                Golang <a href="/api">API</a> running as a <a href="https://vercel.com/docs/serverless-functions">Vercel Serverless Function</a> which returns your public IP address
            </div><br>
            <div style="display: inline-block;">
                API supports the following <code>accept</code> headers:
                <ul>
                    <li><code>application/json</code></li>
                    <li><code>application/protobuf</code></li>
                </ul>
            </div>
        </div>
    </div>
</body>

</html>
`))
