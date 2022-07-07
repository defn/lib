include("/home/ubuntu/Tiltfile")

load("ext://uibutton", "cmd_button", "location")

local_resource(
    "python",
    serve_cmd="""
        cd; cd work/cloud;
        (python -mvenv .v);
        . .v/bin/activate;
        p export src::;
        code --install-extension ms-python.python || true;
        code --install-extension bungcip.better-toml || true;
        p --loop fmt lint check package ::;
    """,
    allow_parallel=True,
    labels=["automation"],
)

cmd_button(
    name="zmake login",
    text="make login",
    icon_name="login",
    argv=[
        "bash",
        "-c",
        """
            cd; cd work/cloud;
            make login;
        """,
    ],
    location="nav",
)

for aname in ["gyre", "curl", "coil", "spiral", "helix"]:
    cmd_button(
        name=aname,
        text=aname,
        icon_name="login",
        argv=[
            "bash",
            "-c",
            """
                xdg-open "https://signin.aws.amazon.com/oauth?Action=logout";
                sleep 1;
                aws-vault login {aname};
            """.format(
                aname=aname
            ),
        ],
        location="nav",
    )
