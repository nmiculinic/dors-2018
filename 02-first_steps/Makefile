hello_world: main.go
	go build -tags netgo -ldflags "-extldflags '-lm -lstdc++ -static'" -o hello_world .

docker_image: hello_world
	docker build -t dors-hello-world .
.PHONY: build

build: hello_world docker_image

.PHONY: build

deploy: build
	kubectl apply -f hello_world_pod.yaml

.PHONY:deploy


