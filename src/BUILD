# type: ignore

python_distribution(
    name="defn",
    dependencies=[
        "3rdparty/python:reqs#cdktf",
        "3rdparty/python:reqs#cdktf-cdktf-provider-aws",
        "3rdparty/python:reqs#cdktf-cdktf-provider-digitalocean",
        "3rdparty/python:reqs#cdktf-cdktf-provider-github",
        "3rdparty/python:reqs#cdktf-cdktf-provider-null",
        "3rdparty/python:reqs#cdktf-cdktf-provider-tfe",
        "provider:defn_cdktf_provider_vault",
        "provider:defn_cdktf_provider_cloudflare",
        "src/defn:cli",
        "src/defn:version",
    ],
    repositories=["@pypi"],
    provides=python_artifact(
        name="defn", version=open("src/defn/VERSION", "r").readline().strip()
    ),
    sdist=True,
    wheel=True,
    entry_points={
        "console_scripts": {
            "defn": "defn.cli:main",
        }
    },
)

defn_sources()
