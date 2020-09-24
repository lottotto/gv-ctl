build:
	env GOOS=linux GOARCH=amd64 GO build -o dist/gvctl-linux
	env GOOS=darwin GOARCH=amd64 GO build -o dist/gvctl-darwin
	env GOOS=windows GOARCH=amd64 GO build -o dist/gvctl-windows
	echo build Success!

docker:
	docker build -t gvctl-alpine:0.2.2 -f docker/Dockerfile .
	docker tag gvctl-alpine:0.2.2 danish9966/gvctl-alpine:0.2.1
	docker push danish9966/gvctl-alpine:0.2.1