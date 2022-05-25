from cdktf import NamedRemoteWorkspace, RemoteBackend, TerraformStack
from cdktf_cdktf_provider_github import GithubProvider
from cdktf_cdktf_provider_null import NullProvider
from cdktf_cdktf_provider_tfe import TfeProvider
from constructs import Construct
from defn_cdktf_provider_boundary.boundary import BoundaryProvider
from defn_cdktf_provider_buildkite.buildkite import BuildkiteProvider
from defn_cdktf_provider_cloudflare.cloudflare import CloudflareProvider
from defn_cdktf_provider_vault.vault import VaultProvider


class DemoStack(TerraformStack):
    """cdktf Stack for demonstration"""

    def __init__(self, scope: Construct, namespace: str):
        super().__init__(scope, namespace)

        NullProvider(self, "null")
        BuildkiteProvider(self, "buildkite", organization="defn", api_token="")
        TfeProvider(self, "tfe")
        GithubProvider(self, "github", organization="defn")
        CloudflareProvider(self, "cloudflare")
        VaultProvider(self, "vault", address="")
        BoundaryProvider(self, "boundary", addr="")

        w = NamedRemoteWorkspace(name="bootstrap")
        RemoteBackend(self, organization="defn", workspaces=w)
