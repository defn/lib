package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	j "github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/instance"
	aws "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"
)

func TheStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, j.String("AWS"), &aws.AwsProviderConfig{
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

	a := TheStack(app, "a")
	cdktf.NewCloudBackend(a, &cdktf.CloudBackendProps{
		Hostname:     j.String("app.terraform.io"),
		Organization: j.String("defn"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("a")),
	})

	b := TheStack(app, "b")
	cdktf.NewCloudBackend(b, &cdktf.CloudBackendProps{
		Hostname:     j.String("app.terraform.io"),
		Organization: j.String("defn"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("b")),
	})

	app.Synth()
}
