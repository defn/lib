include('/home/ubuntu/Tiltfile')

load("ext://uibutton", "cmd_button", "location")

local_resource(
    "python",
    serve_cmd="(python -mvenv .v); . .v/bin/activate; p export src::; code --install-extension ms-python.python || true; code --install-extension bungcip.better-toml || true; p --loop fmt lint check package ::",
    allow_parallel=True,
    labels=["automation"],
)

cmd_button(
    name="make login",
    text="make login",
    icon_name="login",
    argv=[
        "bash",
        "-c",
        """
            make login
        """
    ],
    location="nav",
)
