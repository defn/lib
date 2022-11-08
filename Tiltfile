analytics_settings(False)
allow_k8s_contexts("k3d-control")

load("ext://uibutton", "cmd_button", "location")
load("ext://restart_process", "custom_build_with_restart")

default_registry("169.254.32.1:5000")

for app in ("defn", "defm", "client", "worker"):
    local_resource("%s-go" % (app,), "go build -o dist/image-%s/bin cmd/%s/%s.go; echo done" % (app,app,app), deps=["cmd/%s" % (app,)])

for app in ("defn", "defm", "worker"):
    k8s_yaml("cmd/%s/%s.yaml" % (app,app))

    custom_build_with_restart(
        ref=app,
        command=(
            "./bin/image.sh %s ${EXPECTED_REF}" % (app,)
        ),
        entrypoint="/app/bin",
        deps=["dist/image-%s/bin" % (app,)],
        live_update=[
            sync("dist/image-%s/bin" % (app,), "/app/bin"),
        ],
    )

local_resource("vite", serve_cmd="pnpm install; while true; do turbo dev; sleep 1; done", deps=[".vite-mode"])

cmd_button(
    name="client",
    text="Client",
    icon_name="login",
    argv=[
        "bash", "-c",
        """
            ~/work/cloud/dist/image-client/bin
        """,
    ],
    location=location.NAV,
)
