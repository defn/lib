VERSION --shell-out-anywhere --use-chmod --use-host-command --earthly-version-arg --use-copy-link --use-registry-for-with-docker 0.6

IMPORT ./lib AS lib

ubuntu:
    FROM ubuntu

defn:
    FROM +ubuntu

    ARG image

    RUN mkdir -p app

    COPY dist/cmd.defn/bin app/bin

    ENTRYPOINT ["/app/bin"]

    SAVE IMAGE --push ${image}

defm:
    FROM +ubuntu

    ARG image

    RUN mkdir -p app

    COPY dist/cmd.defm/bin app/bin

    ENTRYPOINT ["/app/bin"]

    SAVE IMAGE --push ${image}

pre-commit:
    FROM ghcr.io/defn/dev
    ARG workdir
    DO lib+PRECOMMIT --workdir=${workdir}

synth:
    FROM ghcr.io/defn/dev
    DO lib+SYNTH
    SAVE ARTIFACT cdktf.out/stacks/gyre/cdk.tf.json AS LOCAL cdktf.out/stacks/gyre/
    SAVE ARTIFACT cdktf.out/stacks/curl/cdk.tf.json AS LOCAL cdktf.out/stacks/curl/
    SAVE ARTIFACT cdktf.out/stacks/coil/cdk.tf.json AS LOCAL cdktf.out/stacks/coil/
    SAVE ARTIFACT cdktf.out/stacks/helix/cdk.tf.json AS LOCAL cdktf.out/stacks/helix/
    SAVE ARTIFACT cdktf.out/stacks/spiral/cdk.tf.json AS LOCAL cdktf.out/stacks/spiral/

init:
    FROM ghcr.io/defn/dev
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
    FROM ghcr.io/defn/dev
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
