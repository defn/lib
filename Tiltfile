update_settings(max_parallel_updates=6)

local_resource("pants", cmd="p --loop fmt lint check package ::", allow_parallel=True)

local_resource("dev", cmd="cd && git pull && make dev", allow_parallel=True)

local_resource(
    "pre-commit",
    cmd="while true; do if docker info; then make pc; break; fi; sleep 1; done",
    allow_parallel=True,
)

local_resource(
    "python",
    cmd="(cd && python -mvenv .venv); (. ~/.venv/bin/activate; p export src::); code --install-extension ms-python.python || true; code --install-extension bungcip.better-toml || true",
    allow_parallel=True,
)
