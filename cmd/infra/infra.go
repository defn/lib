package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	j "github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/instance"
	aws "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"

	"github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/organization"
	tfe "github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/provider"
	"github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/workspace"
)

func defnWorkspacesStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	tfe.NewTfeProvider(stack, j.String("tfe"), &tfe.TfeProviderConfig{
		Hostname: j.String("app.terraform.io"),
	})

	org_aws := organization.NewOrganization(stack, j.String("org-aws"), &organization.OrganizationConfig{
		Name:  j.String("org-aws"),
		Email: j.String("org-aws@defn.sh"),
	})

	workspace.NewWorkspace(stack, j.String("meh-a"), &workspace.WorkspaceConfig{
		Name:         j.String("meh-a"),
		Organization: org_aws.Name(),
	})

	return stack
}

func TheStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, j.String("aws"), &aws.AwsProviderConfig{
		Region: j.String("us-west-1"),
	})

	instance := instance.NewInstance(stack, j.String("compute"), &instance.InstanceConfig{
		Ami:          j.String("ami-01456a894f71116f2"),
		InstanceType: j.String("t2.micro"),
	})

	cdktf.NewTerraformOutput(stack, j.String("public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	workspaces := defnWorkspacesStack(app, "workspaces")
	cdktf.NewCloudBackend(workspaces, &cdktf.CloudBackendProps{
		Hostname:     j.String("app.terraform.io"),
		Organization: j.String("defn"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("workspaces")),
	})

	//	a := TheStack(app, "a")
	//	cdktf.NewCloudBackend(a, &cdktf.CloudBackendProps{
	//		Hostname:     j.String("app.terraform.io"),
	//		Organization: j.String("defn"),
	//		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("a")),
	//	})
	//
	//	b := TheStack(app, "b")
	//	cdktf.NewCloudBackend(b, &cdktf.CloudBackendProps{
	//		Hostname:     j.String("app.terraform.io"),
	//		Organization: j.String("defn"),
	//		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("b")),
	//	})

	app.Synth()
}
