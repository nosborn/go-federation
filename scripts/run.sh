#!/bin/sh

exec docker run \
  --mount "type=bind,source=${PWD}/bin,target=/app/bin" \
  --mount "type=bind,source=${PWD}/data,target=/app/data" \
  --mount "type=bind,source=${PWD}/ssl,target=/app/ssl" \
  --mount "type=tmpfs,dst=/tmp" \
  --publish 23:23 \
  --publish 992:992 \
  --read-only \
  --rm \
  go-federation:latest
