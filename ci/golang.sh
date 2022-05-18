#!/usr/bin/env bash

function main {
    goreleaser build --snapshot --rm-dist
}

main "$@"
