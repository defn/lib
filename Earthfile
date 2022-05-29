VERSION --shell-out-anywhere --use-chmod --use-host-command --earthly-version-arg --use-copy-link 0.6

IMPORT ./lib AS lib

get:
    FROM registry.fly.io/defn:dev-tower
    COPY cdktf.json.get cdktf.json
    RUN ~/bin/e cdktf get
    SAVE ARTIFACT .gen/boundary/* AS LOCAL provider.new/defn_cdktf_provider_boundary/
    SAVE ARTIFACT .gen/vault/* AS LOCAL provider.new/defn_cdktf_provider_vault/
    SAVE ARTIFACT .gen/cloudflare/* AS LOCAL provider.new/defn_cdktf_provider_cloudflare/
    SAVE ARTIFACT .gen/buildkite/* AS LOCAL provider.new/defn_cdktf_provider_buildkite/
    

init:
    FROM registry.fly.io/defn:dev-tower
    ARG stack
    COPY --dir provider src 3rdparty .
    COPY BUILDROOT pants pants.toml .isort.cfg .flake8 .
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants sudo chown ubuntu:ubuntu /home/ubuntu/.cache/pants
    RUN --mount=type=cache,target=/home/ubuntu/.cache/pants ~/bin/e pants package src/defn:main
    DO lib+INIT --stack=${stack}

edit:
    FROM +init
    ARG stack
    RUN --secret TFE_TOKEN --secret TF_TOKEN_app_terraform_io --secret AWS_ACCESS_KEY_ID --secret AWS_SECRET_ACCESS_KEY \
        cd cdktf.out/stacks/${stack} && ~/bin/e terraform import aws_organizations_organization.organization o-something
    RUN --secret TFE_TOKEN --secret TF_TOKEN_app_terraform_io --secret AWS_ACCESS_KEY_ID --secret AWS_SECRET_ACCESS_KEY \
        cd cdktf.out/stacks/${stack} && ~/bin/e terraform import aws_organizations_account.spiral 11111111111

plan:
    FROM +init
    ARG stack
    DO lib+PLAN --stack=${stack}

show:
    FROM +init
    ARG stack
    DO lib+SHOW --stack=${stack}

apply:
    FROM +init
    ARG stack
    DO lib+APPLY --stack=${stack}