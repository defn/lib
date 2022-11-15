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

type Admin struct {
	name  string
	email string
}

type Account struct {
	name     string
	region   string
	prefix   string
	domain   string
	accounts []string
	admins   []*Admin
}

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

func AwsOrganizationStack(scope constructs.Construct, org *Account) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, js(org.name))

	aws.NewAwsProvider(stack,
		js("aws"), &aws.AwsProviderConfig{
			Region: js(org.region),
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
	for _, adm := range org.admins {
		identitystore_user := identitystoreuser.NewIdentitystoreUser(stack,
			js(fmt.Sprintf("admin_sso_user_%s", adm.name)),
			&identitystoreuser.IdentitystoreUserConfig{
				DisplayName: js(adm.name),
				UserName:    js(adm.name),
				Name: &identitystoreuser.IdentitystoreUserName{
					GivenName:  js(adm.name),
					FamilyName: js(adm.name),
				},
				Emails: &identitystoreuser.IdentitystoreUserEmails{
					Primary: jsii.Bool(true),
					Type:    js("work"),
					Value:   js(adm.email),
				},
				IdentityStoreId: js(cdktf.Fn_Element(ssoadmin_instance_isid.Expression(), jsii.Number(0)).(string)),
			})

		identitystoregroupmembership.NewIdentitystoreGroupMembership(stack,
			js(fmt.Sprintf("admin_sso_user_%s_membership", adm.name)),
			&identitystoregroupmembership.IdentitystoreGroupMembershipConfig{
				MemberId:        identitystore_user.UserId(),
				GroupId:         identitystore_group.GroupId(),
				IdentityStoreId: js(cdktf.Fn_Element(ssoadmin_instance_isid.Expression(), jsii.Number(0)).(string)),
			})
	}

	// The master account (named "org") must be imported.
	for _, acct := range org.accounts {
		// Create the organization account
		var organizations_account_config organizationsaccount.OrganizationsAccountConfig

		if acct == org.name {
			// The master organization account can't set
			// iam_user_access_to_billing, role_name
			organizations_account_config = organizationsaccount.OrganizationsAccountConfig{
				Name:  js(acct),
				Email: js(fmt.Sprintf("%s%s@%s", org.prefix, org.name, org.domain)),
				Tags:  &map[string]*string{"ManagedBy": js("Terraform")},
			}
		} else {
			// Organization account
			organizations_account_config = organizationsaccount.OrganizationsAccountConfig{
				Name:                   js(acct),
				Email:                  js(fmt.Sprintf("%s%s+%s@%s", org.prefix, org.name, acct, org.domain)),
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

	defn := Admin{name: "defn", email: "iam@defn.sh"}

	// The infra stacks under management.
	orgs := []Account{
		{
			name:     "gyre",
			region:   "us-east-2",
			prefix:   "aws-",
			domain:   "defn.us",
			accounts: append([]string{"ops"}, []string{"gyre"}...),
			admins:   []*Admin{&defn},
		},
		{
			name:     "curl",
			region:   "us-west-2",
			prefix:   "aws-",
			domain:   "defn.us",
			accounts: append(env_accounts, []string{"curl"}...),
			admins:   []*Admin{&defn},
		},
		{
			name:     "coil",
			region:   "us-east-1",
			prefix:   "aws-",
			domain:   "defn.us",
			accounts: append(env_accounts, []string{"coil"}...),
			admins:   []*Admin{&defn},
		},
		{
			name:     "helix",
			region:   "us-east-2",
			prefix:   "aws-",
			domain:   "defn.sh",
			accounts: append(full_accounts, []string{"helix"}...),
			admins:   []*Admin{&defn},
		},
		{
			name:     "spiral",
			region:   "us-west-2",
			prefix:   "aws-",
			domain:   "defn.us",
			accounts: append(full_accounts, []string{"spiral"}...),
			admins:   []*Admin{&defn},
		},
	}

	for _, acc := range orgs {
		// Create a tfc workspace for each stack
		workspace.NewWorkspace(workspaces, js(acc.name), &workspace.WorkspaceConfig{
			Name:                js(acc.name),
			Organization:        js(tfc_org),
			ExecutionMode:       js("local"),
			FileTriggersEnabled: false,
			QueueAllRuns:        false,
			SpeculativeEnabled:  false,
		})

		// Create the aws organization + accounts stack
		aws_org_stack := AwsOrganizationStack(app, &acc)
		cdktf.NewCloudBackend(aws_org_stack, &cdktf.CloudBackendProps{
			Hostname:     js("app.terraform.io"),
			Organization: js(tfc_org),
			Workspaces:   cdktf.NewNamedCloudWorkspace(js(acc.name)),
		})
	}

	// Emit cdk.tf.json
	app.Synth()
}
