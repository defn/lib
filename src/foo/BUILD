python_sources(name="py")

pex_binary(
    name="textual",
    entry_point="textual.py",
)

experimental_run_shell_command(
    name="build_docs",
    command="find {chroot} -ls; pycco -d html/foo {chroot}/src/foo/*.py",
    dependencies=[":py"],
)
