FROM golang:1.17.2-alpine AS build
WORKDIR /
ENV CGO_ENABLED=0
COPY . .
RUN go mod download
WORKDIR /
RUN go build -o ./kubecon .

FROM scratch AS bin-unix
COPY --from=build /kubecon /kubeconapp

FROM bin-unix AS bin-linux

FROM bin-linux as bin
ENV HOMEDIR="/"
ENTRYPOINT ["/kubeconapp"]
EXPOSE 80/tcp