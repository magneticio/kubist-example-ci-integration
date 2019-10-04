# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /src
RUN cd /src && go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app

EXPOSE 8080
