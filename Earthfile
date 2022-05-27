VERSION --shell-out-anywhere --use-chmod --use-host-command --earthly-version-arg --use-copy-link 0.6

IMPORT github.com/defn/cloud/lib:master AS lib

ARG target=github.com/defn/cloud:master+warm
ARG stack

get:
    FROM registry.fly.io/defn:dev-tower
    COPY cdktf.json.get cdktf.json
    RUN ~/bin/e cdktf get
    SAVE ARTIFACT .gen/boundary/* AS LOCAL provider.new/defn_cdktf_provider_boundary/
    SAVE ARTIFACT .gen/vault/* AS LOCAL provider.new/defn_cdktf_provider_vault/
    SAVE ARTIFACT .gen/cloudflare/* AS LOCAL provider.new/defn_cdktf_provider_cloudflare/
    SAVE ARTIFACT .gen/buildkite/* AS LOCAL provider.new/defn_cdktf_provider_buildkite/

warm:
    FROM registry.fly.io/defn:dev-tower
    COPY --dir provider src 3rdparty .
    COPY BUILDROOT pants pants.toml .isort.cfg .flake8 cdktf.json .
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants sudo chown ubuntu:ubuntu /home/ubuntu/.cache/pants
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants ~/bin/e pants package src/defn:main

plan:
    FROM lib+init --target=${target} --stack=${stack}
    DO lib+PLAN --stack=${stack}

show:
    FROM lib+init --target=${target} --stack=${stack}
    DO lib+SHOW --stack=${stack}

apply:
    FROM lib+init --target=${target} --stack=${stack}
    DO lib+APPLY --stack=${stack}