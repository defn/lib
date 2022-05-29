python_sources(name="py")

experimental_shell_command(
    name="pycco_docs",
    command="pycco -d html *.py",
    tools=["pycco"],
    log_output=True,
    outputs=["html/"],
    dependencies=[":py"],
)

archive(
    name="archive_docs",
    files=[":pycco_docs"],
    format="tar.gz",
    tags=["docs"]
)

pex_binary(
    name="main",
    entry_point="main.py"
)

