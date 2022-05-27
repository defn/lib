VERSION --shell-out-anywhere --use-chmod --use-host-command --earthly-version-arg --use-copy-link 0.6

IMPORT github.com/defn/cloud/lib:master AS lib

ARG target=github.com/defn/cloud:master+warm
ARG stack

warm:
    FROM registry.fly.io/defn:dev-tower
    RUN --no-cache echo '{ "language": "python", "app": "dist/src.defn/main.pex synth" }' > cdktf.json
    COPY --dir provider src 3rdparty .
    COPY BUILDROOT pants pants.toml .isort.cfg .flake8 .
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants sudo chown ubuntu:ubuntu /home/ubuntu/.cache/pants
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants ~/bin/e pants package ::
    RUN ~/bin/e dist/src.defn/main.pex synth
    SAVE ARTIFACT dist/* AS LOCAL dist/
    SAVE ARTIFACT cdktf.out/* AS LOCAL cdktf.out/
    SAVE IMAGE --push registry.fly.io/defn:cloud

get:
    FROM +warm
    COPY cdktf.json .
    RUN ~/bin/e cdktf get
    RUN --no-cache find .gen -ls
    SAVE ARTIFACT .gen/boundary/* AS LOCAL provider/defn_cdktf_provider_boundary/
    SAVE ARTIFACT .gen/vault/* AS LOCAL provider/defn_cdktf_provider_vault/
    SAVE ARTIFACT .gen/cloudflare/* AS LOCAL provider/defn_cdktf_provider_cloudflare/
    SAVE ARTIFACT .gen/buildkite/* AS LOCAL provider/defn_cdktf_provider_buildkite/

plan:
    FROM lib+plan --target=${target} --stack=${stack}
    SAVE ARTIFACT dist/* AS LOCAL dist/
    SAVE ARTIFACT cdktf.out/* AS LOCAL cdktf.out/

apply:
    FROM lib+init --target=${target} --stack=${stack}
    DO lib+APPLY --stack=${stack}
    SAVE ARTIFACT dist/* AS LOCAL dist/
    SAVE ARTIFACT cdktf.out/* AS LOCAL cdktf.out/
