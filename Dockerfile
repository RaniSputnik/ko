FROM golang:1.10

# Install dep - dependency management tool
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy project files
WORKDIR /go/src/github.com/RaniSputnik/ko

# Install dependencies and ko api
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

# TODO split container and builder
# so that we don't need go installed
COPY . .
RUN go install -v .

# Run the API
EXPOSE 8080
CMD ["ko"]