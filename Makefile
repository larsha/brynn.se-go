.PHONY: run gce_deploy_nginx gce_deploy_web

run:
	PORT=3000 go run main.go

# usage: make gce_deploy_nginx version=1.0.0
gce_deploy_nginx:
	docker build -t eu.gcr.io/brynn-145714/brynnse/nginx:latest -t eu.gcr.io/brynn-145714/brynnse/nginx:$(version) -f Dockerfile.nginx .
	gcloud docker -- push eu.gcr.io/brynn-145714/brynnse/nginx:latest
	gcloud docker -- push eu.gcr.io/brynn-145714/brynnse/nginx:$(version)
	docker rmi eu.gcr.io/brynn-145714/brynnse/nginx:latest
	docker rmi eu.gcr.io/brynn-145714/brynnse/nginx:$(version)
	kubectl -n brynnse set image deployment/nginx nginx=eu.gcr.io/brynn-145714/brynnse/nginx:$(version)

gce_deploy_web:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/main .
	curl -o .build/ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
	docker build --build-arg CACHEBUST=$(shell date "+%Y%m%d%H%M%S") -t eu.gcr.io/brynn-145714/brynnse/web:latest -t eu.gcr.io/brynn-145714/brynnse/web:$(version) -f Dockerfile.web .
	rm -rf .build

	gcloud docker -- push eu.gcr.io/brynn-145714/brynnse/web:latest
	gcloud docker -- push eu.gcr.io/brynn-145714/brynnse/web:$(version)
	docker rmi eu.gcr.io/brynn-145714/brynnse/web:latest
	docker rmi eu.gcr.io/brynn-145714/brynnse/web:$(version)
	kubectl -n brynnse set image deployment/web web=eu.gcr.io/brynn-145714/brynnse/web:$(version)
