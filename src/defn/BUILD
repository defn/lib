python_sources(name="py")

pex_binary(
    name="main",
    entry_point="main.py",
)

experimental_run_shell_command(
    name="build_docs",
    command="find {chroot} -ls; pycco -d html/defn {chroot}/src/defn/*.py",
    dependencies=[":py"],
)
