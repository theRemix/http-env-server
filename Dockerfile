FROM golang as build
WORKDIR /srv
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-env-server .

FROM scratch
WORKDIR /
COPY --from=build /srv/http-env-server .
CMD ["/http-env-server"]
