FROM golang:latest

# RUN apk add --update git

ADD . $GOPATH/src/build/


# WORKDIR $GOPATH/src/build/
RUN pwd
RUN ls -alhi
RUN  mkdir -p $GOPATH/src/cmd

# go build -o ./app ./app

# Build the Go app
RUN go build -o bin ./src/build/app/
# RUN ls -alhi bin /src/cmd/app/main.go
RUN ls -alhi src/build/
RUN ls -alhi bin

# COPY . $GOPATH/src/build/
COPY ./app/ $GOPATH/src/cmd/
# VOLUME /usr/bin/proj/src/

RUN ls -alhi src/build/app
RUN ls -alhi src/cmd


# RUN go get -u "github.com/gin-gonic/contrib/cors"
# RUN go get -u "github.com/lib/pq"

# Copy go mod and sum files
# COPY go.mod go.sum ./

# Build the Go app
# RUN go build -o main .

EXPOSE 8888

# CMD ["wget", "https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh"]
# CMD ["chmod", "+x", "wait-for-it.sh"]
# CMD ["./wait-for-it.sh", "db:5432", "--"]

# RUN go build -o ./src/cmd/main.go
# CMD ["./src/build/app"]

# CMD ["go", "run", "./src/cmd/main.go"]

# Command to run the executable
CMD ["./bin/app"]

# docker run -d -p 8080:8080 go-docker

# docker exec -it 72ca2488b353 sh

# docker container run -d --name golan --rm -p 8080:8888 -v ~/dev/secgo/godoc/app/:/go/src/cmd/ web