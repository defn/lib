# type: ignore

defn_sources(name="py", dependencies=[":version"])

file(name="earthfile", source="Earthfile")

resources(name="version", sources=["VERSION"])

experimental_shell_command(
    name="pycco_docs",
    command="pycco -d html *.py",
    tools=["pycco"],
    log_output=True,
    outputs=["html/"],
    dependencies=[":py"],
)

archive(
    name="archive_pycco_docs", files=[":pycco_docs"], format="tar.gz", tags=["docs"]
)

defn_binary(name="cli", entry_point="defn.cli:main")

defn_binary(
    name="bean-client",
    entry_point="bean-client.py",
)

defn_binary(
    name="bean-server",
    entry_point="bean-server.py",
)
