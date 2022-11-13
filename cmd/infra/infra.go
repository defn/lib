package main

import (
	"fmt"

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
	tfc_org := "defn"

	app := cdktf.NewApp(nil)

	workspaces := defnWorkspacesStack(app, "workspaces")
	cdktf.NewCloudBackend(workspaces, &cdktf.CloudBackendProps{
		Hostname:     j.String("app.terraform.io"),
		Organization: j.String(tfc_org),
		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("workspaces")),
	})

	accounts := []string{"test-1", "test-2"}

	for _, acc := range accounts {
		ws_name := fmt.Sprintf("%s-ws", acc)

		workspace.NewWorkspace(workspaces, j.String(ws_name), &workspace.WorkspaceConfig{
			Name:          j.String(ws_name),
			Organization:  j.String(tfc_org),
			ExecutionMode: j.String("local"),
		})

		st := TheStack(app, acc)
		cdktf.NewCloudBackend(st, &cdktf.CloudBackendProps{
			Hostname:     j.String("app.terraform.io"),
			Organization: j.String(tfc_org),
			Workspaces:   cdktf.NewNamedCloudWorkspace(j.String(ws_name)),
		})
	}

	app.Synth()
}
