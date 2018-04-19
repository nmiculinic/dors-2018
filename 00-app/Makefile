hello_world: main.go
	go build -tags netgo -ldflags "-extldflags '-lm -lstdc++ -static'" -o hello_world .

docker_image: hello_world
	eval $$(minikube docker-env)
	docker build -t dors-hello-world .
.PHONY: docker_image

build: hello_world docker_image

.PHONY: build

clean:
	kubectl delete all --all

.PHONY: clean
.FORCE:
