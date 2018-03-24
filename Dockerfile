FROM golang:1.10

# Install dep - dependency management tool
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy project files
WORKDIR /go/src/github.com/RaniSputnik/ko

# Install dependencies and ko api
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

# Build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ko .

FROM scratch
# TODO copy root CA's if needed

# Copy binary
WORKDIR /root/
COPY --from=0 /go/src/github.com/RaniSputnik/ko/ko .

# Run the API
EXPOSE 8080
CMD ["./ko"]