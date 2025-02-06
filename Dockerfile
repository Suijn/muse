FROM golang:1.23.5-bullseye AS base
WORKDIR /app

ARG USER=app
ARG GROUP=app

RUN adduser $USER \
&& addgroup $USER $GROUP

FROM base AS build
COPY --chown=$USER:$GROUP go.mod go.mod
COPY --chown=$USER:$GROUP go.sum go.sum
COPY --chown=$USER:$GROUP users users
COPY --chown=$USER:$GROUP common common
COPY --chown=$USER:$GROUP main.go main.go
COPY --chown=$USER:$GROUP run.sh run.sh
RUN go build -o muse

FROM build AS buildtest
COPY --chown=$USER:$GROUP tests tests

FROM build AS setuser
USER $USER

FROM setuser AS runserver
ENTRYPOINT ["./run.sh"]
CMD ["server"]

FROM buildtest as runtests
ENTRYPOINT ["./run.sh"]
CMD ["test"]
