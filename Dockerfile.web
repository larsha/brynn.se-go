FROM scratch
ADD .build/main /
ADD .build/ca-certificates.crt /etc/ssl/certs/
WORKDIR /

CMD ["/main"]
