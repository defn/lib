include("/home/ubuntu/Tiltfile.common")

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

aws_sso = {
    "coil": "https://d-90674c3cfd.awsapps.com/start#/",
    "curl": "https://d-926760a859.awsapps.com/start#/",
    "gyre": "https://d-9a6716e54a.awsapps.com/start#/",
    "spiral": "https://d-926760b322.awsapps.com/start#/",
    "helix": "https://d-9a6716ffd1.awsapps.com/start#/",
}

aws_icon = {
    "coil": "water",
    "curl": "roundabout_right",
    "gyre": "cyclone",
    "spiral": "route",
    "helix": "all_inclusive",
}

for aname in ["gyre", "curl", "coil", "spiral", "helix"]:
    cmd_button(
        name=aname,
        text=aname,
        icon_name=aws_icon[aname],
        argv=[
            "bash",
            "-c",
            """
                xdg-open "https://signin.aws.amazon.com/oauth?Action=logout";
                aws-vault login {aname};
            """.format(
                aname=aname
            ),
        ],
        location="nav",
    )
    cmd_button(
        name=aname + " sso",
        text=aname + " sso",
        icon_name=aws_icon[aname],
        argv=[
            "bash",
            "-c",
            """
                xdg-open "{sso_url}"
            """.format(
                aname=aname,
                sso_url=aws_sso[aname],
            ),
        ],
        location="nav",
    )
