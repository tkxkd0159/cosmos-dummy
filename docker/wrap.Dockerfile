FROM iignite:latest

# move go.mod file after create chain by ignite-cli if you want to use throwaway type
# docker build -t wrapig -f wrap.Dockerfile ./foochain
COPY go.mod /app/go.mod
RUN go mod download
RUN rm /app/go.mod