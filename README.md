# How it works

1) Install goLang.
2) Install packages `go get github.com/BurntSushi/toml github.com/gorilla/mux`
3) At this directory type `go run app.go`

to test creating new movie use curl
`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:3567/movies?page=3`