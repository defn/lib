VERSION --use-registry-for-with-docker --ci-arg 0.7

IMPORT github.com/defn/dev:0.0.87

dev:
    LOCALLY
    RUN --no-cache ./nix-validate

ci:
    FROM +build
    RUN --no-cache /entrypoint nix-validate

image:
    ARG image
    FROM +build
    SAVE IMAGE --push ${image}

build:
    FROM ghcr.io/defn/dev:latest-nix-empty
    COPY +nix-build/store /nix/store
    COPY +nix-build/app /app

nix-build:
    DO dev+NIX_DIRENV
    SAVE ARTIFACT /nix/store store
    SAVE ARTIFACT /app app
