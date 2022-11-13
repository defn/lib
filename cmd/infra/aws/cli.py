from defn.init import once


""" init must run before cdktf """
import typer


cli = typer.Typer()


@cli.command()
def synth():
    once()

    from cdktf import App, NamedRemoteWorkspace, RemoteBackend

    from defn.aws.stack import AwsOrganizationStack

    app = App()

    full_accounts = ["net", "log", "lib", "ops", "sec", "hub", "pub", "dev", "dmz"]
    env_accounts = ["net", "lib", "hub"]

    stack = AwsOrganizationStack(
        app,
        namespace="spiral",
        org="spiral",
        prefix="aws-",
        domain="defn.us",
        region="us-west-2",
        sso_region="us-west-2",
        accounts=full_accounts,
    )
    RemoteBackend(
        stack, organization="defn", workspaces=NamedRemoteWorkspace(name="spiral")
    )

    stack = AwsOrganizationStack(
        app,
        namespace="helix",
        org="helix",
        prefix="aws-",
        domain="defn.sh",
        region="us-east-2",
        sso_region="us-east-2",
        accounts=full_accounts,
    )
    RemoteBackend(
        stack, organization="defn", workspaces=NamedRemoteWorkspace(name="helix")
    )

    stack = AwsOrganizationStack(
        app,
        namespace="coil",
        org="coil",
        prefix="aws-",
        domain="defn.us",
        region="us-east-1",
        sso_region="us-east-1",
        accounts=env_accounts,
    )
    RemoteBackend(
        stack, organization="defn", workspaces=NamedRemoteWorkspace(name="coil")
    )

    stack = AwsOrganizationStack(
        app,
        namespace="curl",
        org="curl",
        prefix="aws-",
        domain="defn.us",
        region="us-west-1",
        sso_region="us-west-2",
        accounts=env_accounts,
    )
    RemoteBackend(
        stack, organization="defn", workspaces=NamedRemoteWorkspace(name="curl")
    )

    stack = AwsOrganizationStack(
        app,
        namespace="gyre",
        org="gyre",
        prefix="aws-",
        domain="defn.us",
        region="us-east-2",
        sso_region="us-east-2",
        accounts=["ops"],
    )
    RemoteBackend(
        stack, organization="defn", workspaces=NamedRemoteWorkspace(name="gyre")
    )

    app.synth()


@cli.command()
def version():
    import pkgutil

    data = pkgutil.get_data("defn", "VERSION")
    if data is not None:
        print(data.decode("utf-8").strip())


def main():
    cli()
