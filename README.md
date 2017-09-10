# go-fakesgiving
Toy web app allowing people to signup to bring things to Fakesgiving (aka Fake Thanksgiving with friends)

## Dependencies
1. [postgres](https://www.postgresql.org/)
2. [go](https://golang.org/doc/install)
3. [go postgres driver](https://github.com/lib/pq): `go get github.com/lib/pq`

## Install
1. Create a postgres database table called "fakesgiving" (and edit main.go to use your postgres username...need to make a config file)
2. Start up the app: `go run main.go` (see [condegangsta/gin](https://github.com/codegangsta/gin) if you would like to use live reload)
3. Go to localhost:3001 in your browser
