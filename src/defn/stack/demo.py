from cdktf import NamedRemoteWorkspace, RemoteBackend
from constructs import Construct

from defn.aws.stack import AwsOrganizationStack


class DemoStack(AwsOrganizationStack):
    """cdktf Stack for demonstration."""

    def __init__(
        self,
        scope: Construct,
        org: str,
        domain: str,
        region: str,
        sso_region: str,
        accounts,
        prefix: str = "aws-",
    ):
        super().__init__(scope, org, prefix, org, domain, region, sso_region, accounts)

        w = NamedRemoteWorkspace(name=org)
        RemoteBackend(self, organization="defn", workspaces=w)
