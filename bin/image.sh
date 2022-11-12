#!/usr/bin/env bash

set -exu

name="$1"; shift
image="$1"; shift

cd "dist/image-${name}"
git init || true
rsync -ia ../../flake.lock ../../*.nix .

git add -f --intent-to-add flake.lock *.nix bin
git update-index --assume-unchanged flake.lock *.nix bin

n build .#go
sudo rm -rf nix/store
time for a in $(nix-store -qR ./result); do rsync -ia $a nix/store/; done

(
    echo '# syntax=docker/dockerfile:1'
    echo FROM alpine
    echo RUN mkdir -p /app
    for a in nix/store/*/; do
        echo COPY --link "$a" "/$a/"
    done
    echo COPY bin /app/bin

    echo ENTRYPOINT [ '"/app/bin"' ]
    echo "ENV PATH $(for a in nix/store/*/; do echo -n "/$a/bin:"; done)/bin"
) > Dockerfile

time env DOCKER_BUILDKIT=1 docker build -t "${image}" .

docker push "${image}"
