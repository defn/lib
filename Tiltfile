analytics_settings(False)
allow_k8s_contexts("k3d-control")

load("ext://uibutton", "cmd_button", "location")
load("ext://restart_process", "custom_build_with_restart")

default_registry("169.254.32.1:5000")

#local_resource("pants-go", serve_cmd="p --loop package cmd::", deps=["cmd"])
for app in ("defn", "defm", "client", "worker"):
    local_resource("go-%s" % (app,), "go build -o dist/cmd.%s/bin cmd/%s/%s.go" % (app,app,app), deps=["cmd/%s" % (app,)])

for app in ("defn", "defm"):
    k8s_yaml("cmd/%s/%s.yaml" % (app,app))

    custom_build_with_restart(
        ref=app,
        command=(
            "earthly --push --remote-cache=${EXPECTED_REGISTRY}/${EXPECTED_IMAGE}-cache +%s --image=${EXPECTED_REF}" % (app,)
        ),
        entrypoint="/app/bin",
        deps=["dist/cmd.%s" % (app,)],
        live_update=[
            sync("dist/cmd.%s/bin" % (app,), "/app/bin"),
        ],
    )

#local_resource("vite", serve_cmd="while true; do turbo dev; sleep 1; done", deps=[".vite-mode"])
