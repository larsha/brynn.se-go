.PHONY: loadtest deploy_nginx

latest=eu.gcr.io/brynn-145714/brynnse/nginx:latest
image=eu.gcr.io/brynn-145714/brynnse/nginx:$(shell date +%s)

# make loadtest url=https://brynn.se
loadtest:
	docker run -it --rm -p 8089:8089 -v `pwd`/devops/locust:/locust fredriklack/docker-locust \
  	/bin/ash -c "locust -H $(url)"

deploy_nginx:
	docker build \
	-t $(latest) -t $(image) \
	-f Dockerfile.nginx .

	gcloud docker -- push $(latest)
	gcloud docker -- push $(image)

	kubectl -n brynnse set image deployment/nginx nginx=$(image)
