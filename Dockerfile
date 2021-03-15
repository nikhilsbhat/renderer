### Description: Dockerfile for Renderer
FROM golang:alpine3.11

COPY renderer /

# Starting
ENTRYPOINT [ "/renderer" ]