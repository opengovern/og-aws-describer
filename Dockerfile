FROM docker.io/golang:alpine as build
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./build/kaytu-aws-describer /
CMD [ "/kaytu-aws-describer" ]