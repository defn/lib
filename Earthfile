VERSION --shell-out-anywhere --use-chmod --use-host-command --earthly-version-arg --use-copy-link 0.6

IMPORT ./lib AS lib

pre-commit:
    FROM registry.fly.io/defn:dev-tower
    WORKDIR /home/ubuntu/work/cloud
    RUN git init
    COPY .pre-commit-config.yaml .
    RUN --mount=type=cache,target=/home/ubuntu/work/cloud/.cache sudo chown ubuntu:ubuntu /home/ubuntu/work/cloud/.cache
    RUN --mount=type=cache,target=/home/ubuntu/work/cloud/.cache ~/bin/e env PRE_COMMIT_HOME=/home/ubuntu/work/cloud/.cache/pre-commit pre-commit install
    RUN --mount=type=cache,target=/home/ubuntu/work/cloud/.cache ~/bin/e env PRE_COMMIT_HOME=/home/ubuntu/work/cloud/.cache/pre-commit pre-commit run --all
    RUN --mount=type=cache,target=/home/ubuntu/work/cloud/.cache tar cfz pre-commit.tgz .cache
    SAVE ARTIFACT pre-commit.tgz AS LOCAL .cache/pre-commit.tgz

get:
    FROM registry.fly.io/defn:dev-tower
    COPY cdktf.json.get cdktf.json
    RUN ~/bin/e cdktf get
    SAVE ARTIFACT .gen/boundary/* AS LOCAL provider.new/defn_cdktf_provider_boundary/
    SAVE ARTIFACT .gen/vault/* AS LOCAL provider.new/defn_cdktf_provider_vault/
    SAVE ARTIFACT .gen/cloudflare/* AS LOCAL provider.new/defn_cdktf_provider_cloudflare/
    SAVE ARTIFACT .gen/buildkite/* AS LOCAL provider.new/defn_cdktf_provider_buildkite/

synth:
    FROM registry.fly.io/defn:dev-tower
    COPY --dir provider src 3rdparty .
    COPY BUILDROOT pants pants.toml .isort.cfg .flake8 .
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants sudo chown ubuntu:ubuntu /home/ubuntu/.cache/pants
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants ~/bin/e p package src/defn:cli
    DO lib+SYNTH
    SAVE ARTIFACT cdktf.out/stacks/gyre/cdk.tf.json AS LOCAL cdktf.out/stacks/gyre/
    SAVE ARTIFACT cdktf.out/stacks/curl/cdk.tf.json AS LOCAL cdktf.out/stacks/curl/
    SAVE ARTIFACT cdktf.out/stacks/coil/cdk.tf.json AS LOCAL cdktf.out/stacks/coil/
    SAVE ARTIFACT cdktf.out/stacks/helix/cdk.tf.json AS LOCAL cdktf.out/stacks/helix/
    SAVE ARTIFACT cdktf.out/stacks/spiral/cdk.tf.json AS LOCAL cdktf.out/stacks/spiral/

init:
    FROM registry.fly.io/defn:dev-tower
    ARG stack
    DO lib+INIT --stack=${stack}

plan:
    FROM +init
    ARG stack
    DO lib+PLAN --stack=${stack}

show:
    FROM +init
    ARG stack
    DO lib+SHOW --stack=${stack}

import:
    FROM +init
    ARG stack
    DO lib+IMPORT --stack=${stack}

config:
    FROM registry.fly.io/defn:dev-tower
    ARG stack
    ARG region
    ARG sso_region
    ARG sso_url
    DO lib+CONFIG --stack=${stack} --region=${region} --sso_region=${sso_region} --sso_url=${sso_url}

apply:
    FROM +init
    ARG stack
    DO lib+APPLY --stack=${stack}

edit:
    FROM +init
    ARG stack
    ARG cmd
    DO lib+EDIT --stack=${stack} --cmd=${cmd}
