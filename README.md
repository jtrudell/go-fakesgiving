# go-fakesgiving
Toy web app allowing people to signup to bring things to Fakesgiving (aka Fake Thanksgiving with friends)

## Dependencies
1. [postgres](https://www.postgresql.org/)
2. [go](https://golang.org/doc/install)
3. [go postgres driver](https://github.com/lib/pq): `go get github.com/lib/pq`

## Install
1. Edit the .env file in root directory to replace "DBUSER=jentrudell" with "DBUSER=your postgres username"
2. If you would like a random Giphy image for thanksgiving, you need an API key from Giphy. They are free. Once you have it, set it in .env
3. Create a postgres database table called "fakesgiving" (or whatever name you give the database in .env)
4. Start up the app: `go run main.go`
5. Go to localhost:8080 in your browser (or whatever port you set in .env)
