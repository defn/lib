python_sources(name="py")
file(name="earthfile", source="Earthfile")

experimental_shell_command(
    name="pycco_docs",
    command="pycco -d html *.py",
    tools=["pycco"],
    log_output=True,
    outputs=["html/"],
    dependencies=[":py"],
)

experimental_shell_command(
    name="earthly_docs",
    command="find . -ls; DOCKER_HOST=localhost:2375 earthly +build",
    tools=["find", "earthly", "docker", "docker-credential-pass"],
    log_output=True,
    outputs=["dist/"],
    dependencies=[":earthfile", ":pycco_docs"],
)

archive(
    name="archive_pycco_docs", files=[":pycco_docs"], format="tar.gz", tags=["docs"]
)
archive(
    name="archive_earthly_docs", files=[":earthly_docs"], format="tar.gz", tags=["docs"]
)

pex_binary(name="cli", entry_point="cli.py")
