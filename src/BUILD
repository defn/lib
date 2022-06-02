python_distribution( # type: ignore
    name="defn",
    dependencies=[
        "3rdparty/python:reqs#cdktf",
        "3rdparty/python:reqs#cdktf-cdktf-provider-aws",
        "3rdparty/python:reqs#cdktf-cdktf-provider-digitalocean",
        "3rdparty/python:reqs#cdktf-cdktf-provider-github",
        "3rdparty/python:reqs#cdktf-cdktf-provider-null",
        "3rdparty/python:reqs#cdktf-cdktf-provider-tfe",
        "provider:defn_cdktf_provider_boundary",
        "provider:defn_cdktf_provider_vault",
        "provider:defn_cdktf_provider_buildkite",
        "provider:defn_cdktf_provider_cloudflare",
        "src/defn:cli"
    ],
    repositories=["@pypi"],
    provides=python_artifact( # type: ignore
        name="defn",
        version="0.0.1a4"
    ),
    sdist=True,
    wheel=False,
    entry_points={
        "console_scripts": {
            "defn": "defn.cli:main",
        }
    },
)
