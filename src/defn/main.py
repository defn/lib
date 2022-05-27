from foo.init import once


""" init must run before cdktf """

import typer
from cdktf import App

from foo.demo import DemoStack
from foo.textual import GridTest


once()
cli = typer.Typer()


@cli.command()
def synth(name: str = "spiral"):
    app = App()

    DemoStack(
        app,
        namespace=name,
        prefix="aws-",
        org=name,
        domain="defn.us",
        region="us-west-2",
    )

    app.synth()


@cli.command()
def meh():
    GridTest.run(title="Grid Test", log="textual.log")


if __name__ == "__main__":
    cli()
