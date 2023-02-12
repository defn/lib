for cmd in ["hello", "bye", "api", "infra"]:
    local_resource(
        "build-{cmd}".format(cmd=cmd),
        "go build -o go/bin/{cmd} ./cmd/{cmd}".format(cmd=cmd),
        deps=["./cmd/{cmd}".format(cmd=cmd)])
