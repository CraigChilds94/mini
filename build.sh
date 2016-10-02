rm /bin/main
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/main .
docker build .