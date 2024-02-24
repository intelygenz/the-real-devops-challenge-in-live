FROM golang:1.22.0

WORKDIR /app
ADD ./main.go .
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch

WORKDIR /app
COPY --from=0 /app/main .
COPY ./public/index.html public/index.html
COPY ./public/script.js public/script.js
COPY ./public/style.css public/style.css

ENTRYPOINT ["/app/main"]
CMD ["--port", "3000"]

EXPOSE 3000
