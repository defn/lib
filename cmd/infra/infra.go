package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	aws "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"

	tfe "github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/provider"
	"github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/workspace"
)

// alias
func js(s string) *string {
	return jsii.String(s)
}

func defnWorkspacesStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	tfe.NewTfeProvider(stack, js("tfe"), &tfe.TfeProviderConfig{
		Hostname: js("app.terraform.io"),
	})

	return stack
}

func TheStack(scope constructs.Construct, id string, region string, sso_region string, namespace string, org string, prefix string, domain string, sub_accounts []string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, js("aws_sso"), &aws.AwsProviderConfig{
		Region: js(region),
	})

	return stack
}

func main() {
	// Stacks under one tfc organization.
	tfc_org := "defn"

	app := cdktf.NewApp(nil)

	// Bootstrap stack to create workspaces.  Manually create the `workspaces`
	// workspace.
	workspaces := defnWorkspacesStack(app, "workspaces")
	cdktf.NewCloudBackend(workspaces, &cdktf.CloudBackendProps{
		Hostname:     js("app.terraform.io"),
		Organization: js(tfc_org),
		Workspaces:   cdktf.NewNamedCloudWorkspace(js("workspaces")),
	})

	full_accounts := []string{"net", "log", "lib", "ops", "sec", "hub", "pub", "dev", "dmz"}
	env_accounts := []string{"net", "lib", "hub"}

	// The infra stacks under management.
	accounts := []string{"gyre", "curl", "coil", "helix", "spiral"}
	regions := []string{"us-east-2", "us-west-1", "us-east-1", "us-east-2", "us-west-2"}
	sso_regions := []string{"us-east-2", "us-west-2", "us-east-1", "us-east-2", "us-west-2"}
	namespaces := []string{"gyre", "curl", "coil", "helix", "spiral"}
	orgs := []string{"gyre", "curl", "coil", "helix", "spiral"}
	prefixes := []string{"aws-", "aws-", "aws-", "aws-", "aws-"}
	domains := []string{"defn.us", "defn.us", "defn.us", "defn.us", "defn.us"}
	sub_accounts := [][]string{{"ops"}, env_accounts, env_accounts, full_accounts, full_accounts}

	for i, acc := range accounts {
		ws_name := acc

		// Create a tfc workspace for each stack
		workspace.NewWorkspace(workspaces, js(ws_name), &workspace.WorkspaceConfig{
			Name:          js(ws_name),
			Organization:  js(tfc_org),
			ExecutionMode: js("local"),
		})

		// Create the infra stack.
		st := TheStack(app, acc, regions[i], sso_regions[i], namespaces[i], orgs[i], prefixes[i], domains[i], sub_accounts[i])
		cdktf.NewCloudBackend(st, &cdktf.CloudBackendProps{
			Hostname:     js("app.terraform.io"),
			Organization: js(tfc_org),
			Workspaces:   cdktf.NewNamedCloudWorkspace(js(ws_name)),
		})
	}

	// Emit cdk.tf.json
	app.Synth()
}
