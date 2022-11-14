    # The master account (named "org") must be imported.
    for acct in accounts:
        account(
            self,
            prefix,
            org,
            domain,
            acct,
            identitystore_group,
            sso_permission_set_admin,
        )
        """Create the organization account."""
        if acct == org:
            # The master organization account can't set
            # iam_user_access_to_billing, role_name
            organizations_account = OrganizationsAccount(
                self,
                acct,
                name=acct,
                email=f"{prefix}{org}@{domain}",
                tags={"ManagedBy": "Terraform"},
            )
        else:
            # Organization account
            organizations_account = OrganizationsAccount(
                self,
                acct,
                name=acct,
                email=f"{prefix}{org}+{acct}@{domain}",
                iam_user_access_to_billing="ALLOW",
                role_name="OrganizationAccountAccessRole",
                tags={"ManagedBy": "Terraform"},
            )

        # Organization accounts grant Administrator permission set to the Administrator group
        SsoadminAccountAssignment(
            self,
            f"{acct}_admin_sso_account_assignment",
            instance_arn=sso_permission_set_admin.instance_arn,
            permission_set_arn=sso_permission_set_admin.arn,
            principal_id=identitystore_group.group_id,
            principal_type="GROUP",
            target_id=organizations_account.id,
            target_type="AWS_ACCOUNT",
        )
