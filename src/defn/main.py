import foo.init

""" init must run before cdktf """

import typer
from cdktf import App

from foo.demo import DemoStack
from foo.textual import GridTest

cli = typer.Typer()


@cli.command()
def synth(name: str = "demo"):
    app = App()

    stack = DemoStack(app, name)

    app.synth()


@cli.command()
def meh():
    GridTest.run(title="Grid Test", log="textual.log")


if __name__ == "__main__":
    cli()
