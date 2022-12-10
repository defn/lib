SHELL := /bin/bash

menu: # This menu
	@perl -ne 'printf("%20s: %s\n","$$1","$$2") if m{^([\w+-]+):[^#]+#\s(.+)$$}' $(shell ls -d GNUmakefile Makefile* 2>/dev/null)

pub:
	npx vite build
	echo preview > .vite-mode

dev:
	echo dev > .vite-mode

go-update:
	go get -u ./cmd/...
	go mod tidy

release:
	 git tag $$(cat VERSION ); git push origin $$(cat VERSION )
