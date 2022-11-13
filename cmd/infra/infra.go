package main

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/instance"
	aws "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"

	tfe "github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/provider"
	"github.com/cdktf/cdktf-provider-tfe-go/tfe/v3/workspace"
)

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

func TheStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, js("aws"), &aws.AwsProviderConfig{
		Region: js("us-west-1"),
	})

	instance := instance.NewInstance(stack, js("compute"), &instance.InstanceConfig{
		Ami:          js("ami-01456a894f71116f2"),
		InstanceType: js("t2.micro"),
	})

	cdktf.NewTerraformOutput(stack, js("public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})

	return stack
}

func main() {
	tfc_org := "defn"

	app := cdktf.NewApp(nil)

	workspaces := defnWorkspacesStack(app, "workspaces")
	cdktf.NewCloudBackend(workspaces, &cdktf.CloudBackendProps{
		Hostname:     js("app.terraform.io"),
		Organization: js(tfc_org),
		Workspaces:   cdktf.NewNamedCloudWorkspace(js("workspaces")),
	})

	accounts := []string{"test-1", "test-2"}

	for _, acc := range accounts {
		ws_name := fmt.Sprintf("%s-ws", acc)

		workspace.NewWorkspace(workspaces, js(ws_name), &workspace.WorkspaceConfig{
			Name:          js(ws_name),
			Organization:  js(tfc_org),
			ExecutionMode: js("local"),
		})

		st := TheStack(app, acc)
		cdktf.NewCloudBackend(st, &cdktf.CloudBackendProps{
			Hostname:     js("app.terraform.io"),
			Organization: js(tfc_org),
			Workspaces:   cdktf.NewNamedCloudWorkspace(js(ws_name)),
		})
	}

	app.Synth()
}
