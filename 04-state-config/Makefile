build:
	$(MAKE) -C ../00-app build

.PHONY: build

deploy: build
	kubectl apply -f hello_world_stateful.yaml

.PHONY:deploy

clean:
	kubectl delete all --all

.PHONY: clean
