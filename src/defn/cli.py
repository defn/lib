from foo.init import once


""" init must run before cdktf """

import typer
from cdktf import App

from foo.cli import GridTest
from foo.demo import DemoStack


once()
cli = typer.Typer()


@cli.command()
def synth(name: str = "spiral", backup: str = "helix"):
    app = App()

    DemoStack(
        app,
        prefix="aws-",
        org=name,
        domain="defn.us",
        region="us-west-2",
    )

    DemoStack(
        app,
        prefix="aws-",
        org=backup,
        domain="defn.sh",
        region="us-east-2",
    )

    app.synth()


@cli.command()
def meh():
    GridTest.run(title="Grid Test", log="textual.log")


if __name__ == "__main__":
    cli()
