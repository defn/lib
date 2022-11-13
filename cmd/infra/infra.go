package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	j "github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/instance"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v10/provider"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	awsprovider.NewAwsProvider(stack, j.String("AWS"), &awsprovider.AwsProviderConfig{
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

	stack := NewMyStack(app, "default")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendProps{
		Hostname:     j.String("app.terraform.io"),
		Organization: j.String("defn"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(j.String("meh")),
	})

	app.Synth()
}
