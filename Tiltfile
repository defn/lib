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

for app in ("defn", "defm", "client", "worker"):
    local_resource("%s-go" % (app,), "mkdir -p dist/image-%s/app && go build -o dist/image-%s/app/bin cmd/%s/%s.go; echo done" % (app,app,app,app), deps=["cmd/%s" % (app,)])

    if app in ("defn","defm","worker"):
        k8s_yaml("cmd/%s/%s.yaml" % (app,app))

        custom_build_with_restart(
            ref=app,
            command=(
                "c nix-docker-build %s .#go ${EXPECTED_REF}" % (app,)
            ),
            entrypoint="/app/bin",
            deps=["dist/image-%s/app/bin" % (app,)],
            live_update=[
                sync("dist/image-%s/app/bin" % (app,), "/app/bin"),
            ],
        )

cmd_button(
    name="client",
    text="Client",
    icon_name="login",
    argv=[
        "bash", "-c",
        """
            ./dist/image-client/app/bin
        """,
    ],
    location=location.NAV,
)
