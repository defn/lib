package main

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/dataawsssoadmininstances"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/identitystoregroup"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/identitystoregroupmembership"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/identitystoreuser"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/organizationsaccount"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/organizationsorganization"
	aws "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/ssoadminaccountassignment"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/ssoadminmanagedpolicyattachment"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/ssoadminpermissionset"

	tfe "github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/provider"
	"github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/workspace"
)

// alias
func js(s string) *string {
	return jsii.String(s)
}

func TfcOrganizationWorkspacesStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	tfe.NewTfeProvider(stack, js("tfe"), &tfe.TfeProviderConfig{
		Hostname: js("app.terraform.io"),
	})

	return stack
}

func AwsOrganizationStack(scope constructs.Construct, id string, region string, org string, prefix string, domain string, sub_accounts []string, initial_admins []string, admin_emails []string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack,
		js("aws"), &aws.AwsProviderConfig{
			Region: js(region),
		})

	organizationsorganization.NewOrganizationsOrganization(stack,
		js("organization"),
		&organizationsorganization.OrganizationsOrganizationConfig{
			FeatureSet: js("ALL"),
			EnabledPolicyTypes: &[]*string{
				js("SERVICE_CONTROL_POLICY"),
				js("TAG_POLICY")},
			AwsServiceAccessPrincipals: &[]*string{
				js("cloudtrail.amazonaws.com"),
				js("config.amazonaws.com"),
				js("ram.amazonaws.com"),
				js("ssm.amazonaws.com"),
				js("sso.amazonaws.com"),
				js("tagpolicies.tag.amazonaws.com")},
		})
	// Lookup pre-enabled AWS SSO instance
	ssoadmin_instance := dataawsssoadmininstances.NewDataAwsSsoadminInstances(stack,
		js("sso_instance"),
		&dataawsssoadmininstances.DataAwsSsoadminInstancesConfig{})

	ssoadmin_instance_arn := cdktf.NewTerraformLocal(stack,
		js("sso_instance_arn"),
		ssoadmin_instance.Arns())

	ssoadmin_permission_set := ssoadminpermissionset.NewSsoadminPermissionSet(stack,
		js("admin_sso_permission_set"),
		&ssoadminpermissionset.SsoadminPermissionSetConfig{
			Name:            js("Administrator"),
			InstanceArn:     js(cdktf.Fn_Element(ssoadmin_instance_arn.Expression(), jsii.Number(0)).(string)),
			SessionDuration: js("PT2H"),
			Tags:            &map[string]*string{"ManagedBy": js("Terraform")},
		})

	sso_permission_set_admin := ssoadminmanagedpolicyattachment.NewSsoadminManagedPolicyAttachment(stack,
		js("admin_sso_managed_policy_attachment"),
		&ssoadminmanagedpolicyattachment.SsoadminManagedPolicyAttachmentConfig{
			InstanceArn:      ssoadmin_permission_set.InstanceArn(),
			PermissionSetArn: ssoadmin_permission_set.Arn(),
			ManagedPolicyArn: js("arn:aws:iam::aws:policy/AdministratorAccess"),
		})

	ssoadmin_instance_isid := cdktf.NewTerraformLocal(stack,
		js("sso_instance_isid"),
		ssoadmin_instance.IdentityStoreIds())

	// Create Administrators group
	identitystore_group := identitystoregroup.NewIdentitystoreGroup(stack,
		js("administrators_sso_group"),
		&identitystoregroup.IdentitystoreGroupConfig{
			DisplayName:     js("Administrators"),
			IdentityStoreId: js(cdktf.Fn_Element(ssoadmin_instance_isid.Expression(), jsii.Number(0)).(string)),
		})

	// Create initial users in the Administrators group
	for i, name := range initial_admins {
		identitystore_user := identitystoreuser.NewIdentitystoreUser(stack,
			js(fmt.Sprintf("admin_sso_user_%s", name)),
			&identitystoreuser.IdentitystoreUserConfig{
				DisplayName: js(name),
				UserName:    js(name),
				Name: &identitystoreuser.IdentitystoreUserName{
					GivenName:  js(name),
					FamilyName: js(name),
				},
				Emails: &identitystoreuser.IdentitystoreUserEmails{
					Primary: jsii.Bool(true),
					Type:    js("work"),
					Value:   js(admin_emails[i]),
				},
				IdentityStoreId: js(cdktf.Fn_Element(ssoadmin_instance_isid.Expression(), jsii.Number(0)).(string)),
			})

		identitystoregroupmembership.NewIdentitystoreGroupMembership(stack,
			js(fmt.Sprintf("admin_sso_user_%s_membership", name)),
			&identitystoregroupmembership.IdentitystoreGroupMembershipConfig{
				MemberId:        identitystore_user.UserId(),
				GroupId:         identitystore_group.GroupId(),
				IdentityStoreId: js(cdktf.Fn_Element(ssoadmin_instance_isid.Expression(), jsii.Number(0)).(string)),
			})
	}

	// The master account (named "org") must be imported.
	for _, acct := range sub_accounts {
		// Create the organization account
		var organizations_account_config organizationsaccount.OrganizationsAccountConfig

		if acct == org {
			// The master organization account can't set
			// iam_user_access_to_billing, role_name
			organizations_account_config = organizationsaccount.OrganizationsAccountConfig{
				Name:  js(acct),
				Email: js(fmt.Sprintf("%s%s@%s", prefix, org, domain)),
				Tags:  &map[string]*string{"ManagedBy": js("Terraform")},
			}
		} else {
			// Organization account
			organizations_account_config = organizationsaccount.OrganizationsAccountConfig{
				Name:                   js(acct),
				Email:                  js(fmt.Sprintf("%s%s+%s@%s", prefix, org, acct, domain)),
				Tags:                   &map[string]*string{"ManagedBy": js("Terraform")},
				IamUserAccessToBilling: js("ALLOW"),
				RoleName:               js("OrganizationAccountAccessRole"),
			}
		}

		organizations_account := organizationsaccount.NewOrganizationsAccount(stack,
			js(acct),
			&organizations_account_config)

		// Organization accounts grant Administrator permission set to the Administrators group
		ssoadminaccountassignment.NewSsoadminAccountAssignment(stack,
			js(fmt.Sprintf("%s_admin_sso_account_assignment", acct)),
			&ssoadminaccountassignment.SsoadminAccountAssignmentConfig{
				InstanceArn:      sso_permission_set_admin.InstanceArn(),
				PermissionSetArn: sso_permission_set_admin.PermissionSetArn(),
				PrincipalId:      identitystore_group.GroupId(),
				PrincipalType:    js("GROUP"),
				TargetId:         organizations_account.Id(),
				TargetType:       js("AWS_ACCOUNT"),
			})
	}

	return stack
}

func main() {
	// Stacks under one tfc organization.
	tfc_org := "defn"

	app := cdktf.NewApp(nil)

	// Bootstrap stack to create workspaces.  Manually create the `workspaces`
	// workspace.
	workspaces := TfcOrganizationWorkspacesStack(app, "workspaces")
	cdktf.NewCloudBackend(workspaces, &cdktf.CloudBackendProps{
		Hostname:     js("app.terraform.io"),
		Organization: js(tfc_org),
		Workspaces:   cdktf.NewNamedCloudWorkspace(js("workspaces")),
	})

	full_accounts := []string{"net", "log", "lib", "ops", "sec", "hub", "pub", "dev", "dmz"}
	env_accounts := []string{"net", "lib", "hub"}
	defn := []string{"defn"}
	defne := []string{"iam@defn.sh"}

	// The infra stacks under management.
	accounts := []string{"gyre", "curl", "coil", "helix", "spiral"}
	regions := []string{"us-east-2", "us-west-2", "us-east-1", "us-east-2", "us-west-2"}
	namespaces := []string{"gyre", "curl", "coil", "helix", "spiral"}
	orgs := []string{"gyre", "curl", "coil", "helix", "spiral"}
	prefixes := []string{"aws-", "aws-", "aws-", "aws-", "aws-"}
	domains := []string{"defn.us", "defn.us", "defn.us", "defn.sh", "defn.us"}
	sub_accounts := [][]string{{"ops"}, env_accounts, env_accounts, full_accounts, full_accounts}
	initial_admins := [][]string{defn, defn, defn, defn, defn}
	admin_emails := [][]string{defne, defne, defne, defne, defne}

	for i, acc := range accounts {
		// Create a tfc workspace for each stack
		workspace.NewWorkspace(workspaces, js(acc), &workspace.WorkspaceConfig{
			Name:                js(acc),
			Organization:        js(tfc_org),
			ExecutionMode:       js("local"),
			FileTriggersEnabled: false,
			QueueAllRuns:        false,
			SpeculativeEnabled:  false,
		})

		// Create the aws organization + accounts stack
		aws_org_stack := AwsOrganizationStack(app, namespaces[i], regions[i], orgs[i], prefixes[i], domains[i], append([]string{orgs[i]}, sub_accounts[i]...), initial_admins[i], admin_emails[i])
		cdktf.NewCloudBackend(aws_org_stack, &cdktf.CloudBackendProps{
			Hostname:     js("app.terraform.io"),
			Organization: js(tfc_org),
			Workspaces:   cdktf.NewNamedCloudWorkspace(js(acc)),
		})
	}

	// Emit cdk.tf.json
	app.Synth()
}
