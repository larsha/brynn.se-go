.PHONY: loadtest

# make loadtest url=https://brynn.se
loadtest:
	docker run -it --rm -p 8089:8089 -v `pwd`/devops/locust:/locust fredriklack/docker-locust \
  	/bin/ash -c "locust -H $(url)"
