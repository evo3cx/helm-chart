build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/helmChart main.go

buildDocker:
	make build /
	&& docker build -t helmChart .
