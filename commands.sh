#!/usr/bin/env bash
set -ex
build() {
  pushd apps/list_players
  go install
  popd
}

run() {
  build
  list_players
}

"$@"
