.PHONY: run scratch gce_deploy

run:
	PORT=3000 go run main.go

scratch:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main .

	curl -o .build/ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
	docker build -f Dockerfile.web -t brynn.se-go .
	rm -rf .build

	docker run --rm -it \
		-p 3000:3000 \
		brynn.se-go

# usage: make gce_deploy service=web version=1.0.0
gce_deploy:
	docker build -t eu.gcr.io/brynn-145714/brynn.se-go/$(service):latest -t eu.gcr.io/brynn-145714/brynn.se-go/$(service):$(version) -f Dockerfile.web .
	gcloud docker -- push eu.gcr.io/brynn-145714/brynn.se-go/$(service):latest
	gcloud docker -- push eu.gcr.io/brynn-145714/brynn.se-go/$(service):$(version)
	docker rmi eu.gcr.io/brynn-145714/brynn.se-go/$(service):latest
	docker rmi eu.gcr.io/brynn-145714/brynn.se-go/$(service):$(version)
	kubectl -n bolagetio set image deployment/$(service) $(service)=eu.gcr.io/brynn-145714/brynn.se-go/$(service):$(version)
