REGISTRY = hub.mtzanidakis.com
IMAGE = dirsizer
VERSION = 0.0.1

.PHONY: all
all: container-push

.PHONY: container
container:
	docker build -t $(REGISTRY)/$(IMAGE):$(VERSION) .
	docker tag $(REGISTRY)/$(IMAGE):$(VERSION) $(REGISTRY)/$(IMAGE):latest

.PHONY: container-push
container-push: container
	docker push $(REGISTRY)/$(IMAGE):$(VERSION)
	docker push $(REGISTRY)/$(IMAGE):latest
