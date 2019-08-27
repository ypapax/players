ARG GO_VERSION=1.12
FROM golang:${GO_VERSION}
COPY . /list_players/
WORKDIR /list_players/apps/list_players
RUN go install
WORKDIR /list_players
CMD list_players -conf $CONF
