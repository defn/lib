from cdktf import NamedRemoteWorkspace, RemoteBackend
from constructs import Construct

from foo.aws import AwsOrganizationStack


class DemoStack(AwsOrganizationStack):
    """cdktf Stack for demonstration."""

    def __init__(
        self,
        scope: Construct,
        namespace: str,
        prefix: str,
        org: str,
        domain: str,
        region: str,
    ):
        super().__init__(scope, namespace, prefix, org, domain, region)

        w = NamedRemoteWorkspace(name="bootstrap")
        RemoteBackend(self, organization="defn", workspaces=w)
