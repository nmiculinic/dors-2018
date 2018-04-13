build:
	$(MAKE) -C ../00-first_steps build

.PHONY: build

deploy: build
	kubectl apply -f hello_world_pod.yaml

.PHONY:deploy

.FORCE:
