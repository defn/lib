VERSION --shell-out-anywhere --use-chmod --use-host-command --earthly-version-arg --use-copy-link 0.6

ARG arch

IMPORT ./lib AS lib

pre-commit:
    FROM defn/dev:${arch}
    ARG workdir
    DO lib+PRECOMMIT --workdir=${workdir}

get:
    FROM defn/dev:${arch}
    COPY cdktf.json.get cdktf.json
    RUN ~/bin/e cdktf get
    SAVE ARTIFACT .gen/boundary/* AS LOCAL provider.new/defn_cdktf_provider_boundary/
    SAVE ARTIFACT .gen/vault/* AS LOCAL provider.new/defn_cdktf_provider_vault/
    SAVE ARTIFACT .gen/cloudflare/* AS LOCAL provider.new/defn_cdktf_provider_cloudflare/
    SAVE ARTIFACT .gen/buildkite/* AS LOCAL provider.new/defn_cdktf_provider_buildkite/

synth:
    FROM defn/dev:${arch}
    RUN ~/bin/e python -mvenv .v
    COPY --dir provider src 3rdparty pants-plugins .
    COPY BUILDROOT pants pants.toml .isort.cfg .flake8 .
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants sudo chown ubuntu:ubuntu /home/ubuntu/.cache/pants
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants . .v/bin/activate && ~/bin/e p package src/defn:cli
    DO lib+SYNTH
    SAVE ARTIFACT cdktf.out/stacks/gyre/cdk.tf.json AS LOCAL cdktf.out/stacks/gyre/
    SAVE ARTIFACT cdktf.out/stacks/curl/cdk.tf.json AS LOCAL cdktf.out/stacks/curl/
    SAVE ARTIFACT cdktf.out/stacks/coil/cdk.tf.json AS LOCAL cdktf.out/stacks/coil/
    SAVE ARTIFACT cdktf.out/stacks/helix/cdk.tf.json AS LOCAL cdktf.out/stacks/helix/
    SAVE ARTIFACT cdktf.out/stacks/spiral/cdk.tf.json AS LOCAL cdktf.out/stacks/spiral/

init:
    FROM defn/dev:${arch}
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
    FROM defn/dev:${arch}
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
