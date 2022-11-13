from cdktf import Fn
from cdktf_cdktf_provider_aws.ssoadmin import (
    SsoadminManagedPolicyAttachment,
    SsoadminPermissionSet,
)


""" Creates Organizations, Accounts, and Administrator permission set """


def administrator(self, ssoadmin_instances):
    """Administrator SSO permission set with AdministratorAccess policy."""
    resource = SsoadminPermissionSet(
        self,
        "admin_sso_permission_set",
        name="Administrator",
        instance_arn=Fn.element(ssoadmin_instances.arns, 0),
        session_duration="PT2H",
        tags={"ManagedBy": "Terraform"},
    )

    SsoadminManagedPolicyAttachment(
        self,
        "admin_sso_managed_policy_attachment",
        instance_arn=resource.instance_arn,
        permission_set_arn=resource.arn,
        managed_policy_arn="arn:aws:iam::aws:policy/AdministratorAccess",
    )

    return resource
