### Description: Dockerfile for Renderer
FROM alpine:3.11

COPY renderer /

# Starting
ENTRYPOINT [ "/renderer" ]