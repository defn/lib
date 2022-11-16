package main

import (
	_ "embed"

	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"

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

//go:embed schema/aws.cue
var aws_schema_cue string

type TerraformCloud struct {
	organization string
	workspace    string
}

type AwsAdmin struct {
	name  string
	email string
}

type AwsOrganization struct {
	name     string
	region   string
	prefix   string
	domain   string
	accounts []string
	admins   []AwsAdmin
}

type AwsProps struct {
	terraform     TerraformCloud
	organizations map[string]AwsOrganization
}

// alias
func js(s string) *string {
	return jsii.String(s)
}

func TfcOrganizationWorkspacesStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, js(id))

	tfe.NewTfeProvider(stack, js("tfe"), &tfe.TfeProviderConfig{
		Hostname: js("app.terraform.io"),
	})

	return stack
}

func AwsOrganizationStack(scope constructs.Construct, org *AwsOrganization) cdktf.TerraformStack {
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
	for _, acct := range append(org.accounts, []string{org.name}...) {
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

func CueToAwsProps(c cue.Value) *AwsProps {
	var props AwsProps

	props.terraform.organization, _ = c.LookupPath(cue.ParsePath("terraform.organization")).Value().String()
	props.terraform.workspace, _ = c.LookupPath(cue.ParsePath("terraform.workspace")).Value().String()

	orgs := make(map[string]AwsOrganization)
	c.LookupPath(cue.ParsePath("organizations")).Value().Decode(&orgs)

	props.organizations = make(map[string]AwsOrganization)

	for name, org := range orgs {
		o := c.LookupPath(cue.ParsePath("organizations." + name))

		org.name = name

		org.region, _ = o.LookupPath(cue.ParsePath("region")).Value().String()
		org.prefix, _ = o.LookupPath(cue.ParsePath("prefix")).Value().String()
		org.domain, _ = o.LookupPath(cue.ParsePath("domain")).Value().String()

		acct_len, _ := o.LookupPath(cue.ParsePath("accounts")).Value().Len().Int64()
		org.accounts = make([]string, acct_len)
		for i := 0; i < int(acct_len); i++ {
			org.accounts[i], _ = o.LookupPath(cue.ParsePath(fmt.Sprintf("accounts[%d]",i))).Value().String()
		}

		admin_len, _ := o.LookupPath(cue.ParsePath("admins")).Value().Len().Int64()
		org.admins = make([]AwsAdmin, admin_len)
		for i := 0; i < int(admin_len); i++ {
			org.admins[i].name, _ = o.LookupPath(cue.ParsePath(fmt.Sprintf("admins[%d].name",i))).Value().String()
			org.admins[i].email, _ = o.LookupPath(cue.ParsePath(fmt.Sprintf("admins[%d].email",i))).Value().String()
		}

		props.organizations[name] = org
	}

	return &props
}

func main() {
	ctx := cuecontext.New()

	user_schema := ctx.CompileString(aws_schema_cue)

	user_input_instance := load.Instances([]string{"."}, nil)[0]
	user_input := ctx.BuildInstance(user_input_instance)

	user_schema.Unify(user_input)

	aws_props := CueToAwsProps(user_input.LookupPath(cue.ParsePath("input")))

	// Our app manages the tfc workspaces, aws organizations plus their accounts
	app := cdktf.NewApp(nil)

	workspaces := TfcOrganizationWorkspacesStack(app, aws_props.terraform.workspace)
	cdktf.NewCloudBackend(workspaces, &cdktf.CloudBackendProps{
		Hostname:     js("app.terraform.io"),
		Organization: js(aws_props.terraform.organization),
		Workspaces:   cdktf.NewNamedCloudWorkspace(js("workspaces")),
	})

	for _, org := range aws_props.organizations {
		// Create a tfc workspace for each stack
		workspace.NewWorkspace(workspaces, js(org.name), &workspace.WorkspaceConfig{
			Name:                js(org.name),
			Organization:        js(aws_props.terraform.organization),
			ExecutionMode:       js("local"),
			FileTriggersEnabled: false,
			QueueAllRuns:        false,
			SpeculativeEnabled:  false,
		})

		// Create the aws organization + accounts stack
		aws_org_stack := AwsOrganizationStack(app, &org)
		cdktf.NewCloudBackend(aws_org_stack, &cdktf.CloudBackendProps{
			Hostname:     js("app.terraform.io"),
			Organization: js(aws_props.terraform.organization),
			Workspaces:   cdktf.NewNamedCloudWorkspace(js(org.name)),
		})
	}

	// Emit cdk.tf.json
	app.Synth()
}
