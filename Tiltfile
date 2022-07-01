analytics_settings(False)

update_settings(max_parallel_updates=6)

local_resource(
    "python",
    serve_cmd="(python -mvenv .v); . .v/bin/activate; p export src::; code --install-extension ms-python.python || true; code --install-extension bungcip.better-toml || true; p --loop fmt lint check package ::",
    allow_parallel=True,
)
