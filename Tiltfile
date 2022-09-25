include("/home/ubuntu/Tiltfile")

load("ext://uibutton", "cmd_button", "location")
load("ext://restart_process", "custom_build_with_restart")

default_registry("169.254.32.1:5000")

local_resource("vite", serve_cmd="while true; do ./node_modules/.bin/vite --host 0.0.0.0; sleep 1; done")

#local_resource("go-defn", "go build -o dist/cmd.defn/bin cmd/defn/main.go", deps=["cmd/defn", "generated"])
#local_resource("go-defm", "go build -o dist/cmd.defm/bin cmd/defm/main.go", deps=["cmd/defm", "generated"])
#
#k8s_yaml("meh.yaml")
#
#custom_build_with_restart(
#    ref="meh",
#    command=(
#        "earthly --push --remote-cache=${EXPECTED_REGISTRY}/${EXPECTED_IMAGE}-cache +meh --image=${EXPECTED_REF}"
#    ),
#    entrypoint="/app/cmd.defn/bin",
#    deps=["dist"],
#    live_update=[
#        sync("dist", "/app"),
#    ],
#)
#