#!/usr/bin/env bash

set -exu

name="$1"; shift
image="$1"; shift

mkdir -p "dist/image-${name}"
cd "dist/image-${name}"
git init || true
rm -f flake.lock flake.nix
cp ../../flake.{nix,lock} .
git add -f --intent-to-add bin flake.nix flake.lock
git update-index --assume-unchanged bin flake.nix flake.lock
mkdir -p nix/store
n build
time for a in $(nix-store -qR ./result); do rsync -ia $a nix/store/; done
(echo '# syntax=docker/dockerfile:1'; echo FROM alpine; echo RUN mkdir -p /app; for a in nix/store/*/; do echo COPY --link "$a" "/$a/"; done; echo COPY bin /app/bin; echo ENTRYPOINT [ '"/app/bin"' ]) > Dockerfile
echo "ENV PATH $(for a in nix/store/*/; do echo -n "/$a/bin:"; done)/bin" >> Dockerfile
time env DOCKER_BUILDKIT=1 docker build -t "${image}" .
docker push "${image}"
