from cdktf import NamedRemoteWorkspace, RemoteBackend
from constructs import Construct

from foo.aws import AwsOrganizationStack


class DemoStack(AwsOrganizationStack):
    """cdktf Stack for demonstration."""

    def __init__(
        self,
        scope: Construct,
        prefix: str,
        org: str,
        domain: str,
        region: str,
    ):
        super().__init__(scope, org, prefix, org, domain, region)

        w = NamedRemoteWorkspace(name=org)
        RemoteBackend(self, organization="defn", workspaces=w)
