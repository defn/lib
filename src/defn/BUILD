# type: ignore

python_sources(name="py", dependencies=[":version"])

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

pex_binary(name="cli", entry_point="defn.cli:main")
