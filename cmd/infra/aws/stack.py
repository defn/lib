from cdktf import TerraformStack
from cdktf_cdktf_provider_aws import AwsProvider
from constructs import Construct

import defn.aws.organizations


""" Creates Organizations, Accounts, and Administrator permission set """


class AwsOrganizationStack(TerraformStack):
    """cdktf Stack for an organization with accounts, sso."""

    def __init__(
        self,
        scope: Construct,
        namespace: str,
        prefix: str,
        org: str,
        domain: str,
        region: str,
        sso_region: str,
        accounts,
    ):
        super().__init__(scope, namespace)

        AwsProvider(self, "aws_sso", region=sso_region)

        defn.aws.organizations.organization(self, prefix, org, domain, [org] + accounts)
