FROM golang:1.24.1-alpine

WORKDIR /tmp/kuma-archive

COPY . .

RUN curl -fsSL https://bun.sh/install | sh

RUN bun install
RUN bun run package

WORKDIR /usr/local
RUN mv /tmp/kuma-archive/dist/generated/*.tar.gz /usr/local
RUN tar -zxvf kuma-archive*.tar.gz
RUN rm -rf kuma-archive*.tar.gz /tmp/kuma-archive

WORKDIR /usr/local/kuma-archive

EXPOSE 8080
ENTRYPOINT [ "/usr/local/kuma-archive", "daemon" ]
