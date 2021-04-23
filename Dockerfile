FROM golang:1.16.3-alpine3.13
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

# EXPOSE 8000
# 
# CMD ["/app/main"]