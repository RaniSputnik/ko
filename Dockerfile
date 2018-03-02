FROM golang:1.10

# Install dep - dependency management tool
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy project files
WORKDIR /go/src/github.com/RaniSputnik/ko
COPY . .

# Install dependencies and ko api
RUN dep ensure
RUN go install -v .

# Run the API
EXPOSE 8080
CMD ["ko"]