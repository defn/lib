analytics_settings(False)

load("ext://uibutton", "cmd_button", "location")

local_resource("vite",
    serve_cmd=[
        "bash", "-c",
        """
            pnpm install
            exec turbo dev
        """
    ],
    deps=[".vite-mode"]
)

local_resource("temporal",
    serve_cmd=[
        "bash", "-c",
        """
            pkill -9 temporalit[e] || true
            rm -f ~/.config/temporalite/db/default.db
            exec temporalite start --namespace default --ip 0.0.0.0
        """
    ]
)

cmd_button(
    name="client",
    text="Client",
    icon_name="login",
    argv=[
        "bash", "-c",
        """
            cd dist/infra/app && ./bin localhost:7233 queue
        """,
    ],
    location=location.NAV,
)

# TODO when infra resource is updated, run infra-test
local_resource("infra-submit",
    deps=[
            "cmd/%s/main.cue" % ("infra",),
        ],
    cmd=[
        "bash", "-c",
        """
            cd cmd/%s && ../../dist/%s/app/bin localhost:7233 queue
        """ % ("infra","infra")
    ]
)

local_resource("infra-workflow",
    deps=[
            "dist/%s/app/bin" % ("infra",),
        ],
    serve_cmd=[
        "bash", "-c",
        """
            cd dist/%s/app && exec ./bin localhost:7233
        """ % ("infra",)
    ]
)

for app in ("api", "infra"):
    local_resource("%s-build" % (app,),
        "mkdir -p dist/%s/app; cp cmd/%s/*.cue dist/%s/app/; mkdir -p dist/%s/app && go build -o dist/%s/app/bin cmd/%s/%s.go; echo done" % (app,app,app,app,app,app,app),
        deps=[
            "cmd/%s/%s.go" % (app,app),
            "cmd/%s/schema/" % (app,)
        ])
