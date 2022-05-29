from foo.init import once


""" init must run before cdktf """

import typer
from cdktf import App

from foo.cli import GridTest
from foo.demo import DemoStack


once()
cli = typer.Typer()


@cli.command()
def synth():
    app = App()

    DemoStack(
        app,
        prefix="aws-",
        org="spiral",
        domain="defn.us",
        region="us-west-2",
        account=["net", "log", "lib", "ops", "sec", "hub", "pub", "dev", "dmz"],
    )

    DemoStack(
        app,
        prefix="aws-",
        org="helix",
        domain="defn.sh",
        region="us-east-2",
        account=["net", "log", "lib", "ops", "sec", "hub", "pub", "dev", "dmz"],
    )

    DemoStack(
        app,
        prefix="aws-",
        org="coil",
        domain="defn.us",
        region="us-east-1",
        account=["net", "lib", "hub"],
    )

    DemoStack(
        app,
        prefix="aws-",
        org="curl",
        domain="defn.us",
        region="us-west-1",
        account=["net", "lib", "hub"],
    )

    app.synth()


@cli.command()
def meh():
    GridTest.run(title="Grid Test", log="textual.log")


if __name__ == "__main__":
    cli()
