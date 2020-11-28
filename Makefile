
gowebexec: main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@

docker: gowebexec
	docker build -t gowebexec:latest .
