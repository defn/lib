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

    full_accounts = (["net", "log", "lib", "ops", "sec", "hub", "pub", "dev", "dmz"],)
    env_accounts = (["net", "lib", "hub"],)

    DemoStack(
        app,
        org="spiral",
        domain="defn.us",
        region="us-west-2",
        account=full_accounts,
    )

    DemoStack(
        app,
        org="helix",
        domain="defn.sh",
        region="us-east-2",
        account=full_accounts,
    )

    DemoStack(
        app,
        org="coil",
        domain="defn.us",
        region="us-east-1",
        account=env_accounts,
    )

    DemoStack(
        app,
        org="curl",
        domain="defn.us",
        region="us-west-1",
        account=env_accounts,
    )

    app.synth()


@cli.command()
def meh():
    GridTest.run(title="Grid Test", log="textual.log")


if __name__ == "__main__":
    cli()
