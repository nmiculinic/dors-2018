build:
	$(MAKE) -C ../02-first_steps build

.PHONY: build

deploy: build
	kubectl apply -f hello_world_k8s.yaml

.PHONY:deploy


