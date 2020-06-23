# First stage: build the executable.
FROM golang:latest AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

# Build the executable to `/app`. Mark the build as statically linked.
RUN go build -installsuffix 'static' -o dating-api
RUN chmod +x dating-api

# Final stage: the running container.
FROM alpine:latest AS final

# Import the compiled executable from the first stage.
COPY --from=builder /src/dating-api /app/dating-api

# Declare the port on which the webserver will be exposed.
EXPOSE 8000

CMD ["/app/dating-api"]
