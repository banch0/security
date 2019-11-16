FROM golang:latest

# environment variable
ENV PORT 8888
# create dir
RUN  mkdir -p $GOPATH/src/cmd

# copy src to container
COPY ./app/ $GOPATH/src/cmd/

# Build the Go app
RUN go build -o bin/app ./src/cmd/

# Mount dir inside
VOLUME $GOPATH/src/

# RUN go get -u "github.com/gin-gonic/contrib/cors"
# RUN go get -u "github.com/lib/pq"

# Copy go mod and sum files
# COPY go.mod go.sum ./

# OPEN PORT
EXPOSE 8888

# Command to run the executable
CMD ["./bin/app"]

# COMMAND FOR RUNNING
# docker build -t web .

# docker exec -it 72ca2488b353 sh

# docker container run -d --name golan --rm -p 8080:8888 -v ~/dev/secgo/godoc/app/:/go/src/cmd/ web


# sudo chown -R $USER:$USER .