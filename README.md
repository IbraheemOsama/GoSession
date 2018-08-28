# How it works

1) Install goLang.
2) Set GOPATH user variable in your OS. (Normally you've a workspace folder that setted in the GOPATH env variable and this folder has a src folder that contains the GoSession repo folder, ex workspace/src/GoSession)
3) Install packages `go get github.com/BurntSushi/toml github.com/gorilla/mux`
4) Install VS Code or
5) Fix the TODOs
6) At this directory type `go run app.go`

to test creating new movie use curl
`curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:3567/movies?page=3`