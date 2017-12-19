.PHONY: loadtest deploy_nginx

latest=eu.gcr.io/brynn-145714/frela/nginx:latest
image=eu.gcr.io/brynn-145714/frela/nginx:$(shell date +%s)

# make loadtest url=https://fre.la
loadtest:
	docker run -it --rm -p 8089:8089 -v `pwd`/devops/locust:/locust fredriklack/docker-locust \
  	/bin/ash -c "locust -H $(url)"

deploy_nginx:
	docker build \
	-t $(latest) -t $(image) \
	-f Dockerfile.nginx .

	gcloud docker -- push $(latest)
	gcloud docker -- push $(image)

	kubectl -n frela set image ds/nginx nginx=$(image)
