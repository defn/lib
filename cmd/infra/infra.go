package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	j "github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/instance"
	aws "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"

	tfe "github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/provider"
	"github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/workspace"
)

func defnWorkspacesStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	tfe.NewTfeProvider(stack, j.String("tfe"), &tfe.TfeProviderConfig{
		Hostname: j.String("app.terraform.io"),
	})

	workspace.NewWorkspace(stack, j.String("test-1-ws"), &workspace.WorkspaceConfig{
		Name:          j.String("test-1-ws"),
		Organization:  j.String("defn"),
		ExecutionMode: j.String("local"),
	})

	workspace.NewWorkspace(stack, j.String("test-2-ws"), &workspace.WorkspaceConfig{
		Name:          j.String("test-2-ws"),
		Organization:  j.String("defn"),
		ExecutionMode: j.String("local"),
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

	test_1 := TheStack(app, "test-1")
	cdktf.NewCloudBackend(test_1, &cdktf.CloudBackendProps{
		Hostname:     j.String("app.terraform.io"),
		Organization: j.String("defn"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("test-1-ws")),
	})

	test_2 := TheStack(app, "test-2")
	cdktf.NewCloudBackend(test_2, &cdktf.CloudBackendProps{
		Hostname:     j.String("app.terraform.io"),
		Organization: j.String("defn"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("test-2-ws")),
	})

	app.Synth()
}
