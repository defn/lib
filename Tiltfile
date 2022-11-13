analytics_settings(False)
allow_k8s_contexts("k3d-control")

load("ext://uibutton", "cmd_button", "location")
load("ext://restart_process", "custom_build_with_restart")

default_registry("169.254.32.1:5000")

local_resource("vite", serve_cmd="pnpm install; while true; do turbo dev; sleep 1; done", deps=[".vite-mode"])

local_resource("temporal",
    serve_cmd=[
        "bash", "-c",
        """
            set -x;
            while true; do
                pkill -9 temporalit[e]
                rm -f ~/.config/temporalite/db/default.db
                temporalite start --namespace default --ip 0.0.0.0
                sleep 10
            done
        """
    ]
)

for app in ("defn", "api", "client", "workflow", "infra"):
    local_resource("%s-go" % (app,),
        "mkdir -p dist/%s/app && go build -o dist/%s/app/bin cmd/%s/%s.go; echo done" % (app,app,app,app),
        deps=["cmd/%s/%s.go" % (app,app)])

    if app in ("defn","api","workflow"):
        k8s_yaml("cmd/%s/%s.yaml" % (app,app))

        custom_build_with_restart(
            ref=app,
            command=(
                "c nix-docker-build %s .#go ${EXPECTED_REF}" % (app,)
            ),
            entrypoint="/app/bin",
            deps=["dist/%s/app/bin" % (app,)],
            live_update=[
                sync("dist/%s/app/bin" % (app,), "/app/bin"),
            ],
        )

    if app in ("infra",):
        local_resource("%s-tf" % (app,),
            deps=["dist/%s/app/bin" % (app,)],
            cmd=[
                "bash", "-c",
                """
                    set -exfu
                    export CDKTF_CONTEXT_JSON="$(jq -n '{excludeStackIdFromLogicalIds: "true", allowSepCharsInLogicalIds: "true"}')"
                    (cd dist/%s && rm -rf cdktf.out && ./app/bin)
                    mkdir -p cmd/%s/tf
                    (set +f; rsync -ia dist/%s/cdktf.out/stacks/. cmd/%s/tf/.)
                    set +x
                    for a in {1..10}; do echo; done
                    set -x
                    git diff cmd/%s/tf || true
                    echo done
                """ % (app,app,app,app,app)
            ]
        )
        local_resource("%s-plan" % (app,),
            deps=["cmd/%s/tf/workspaces/cdk.tf.json" % (app,)],
            cmd=[
                "bash", "-c",
                """
                    set -exfu
                    (cd cmd/%s/tf/workspaces && make plan)
                """ % (app,)
            ]
        )


cmd_button(
    name="client",
    text="Client",
    icon_name="login",
    argv=[
        "bash", "-c",
        """
            ./dist/client/app/bin
        """,
    ],
    location=location.NAV,
)
