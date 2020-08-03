FROM golang:1.14.1
WORKDIR /app
ENV MONGO_HOST mongodb://mongodb:27017
ENV DB_NAME people
COPY . /app
RUN go mod tidy
RUN go build -o main main.go
ENTRYPOINT ./main