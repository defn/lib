package main

import (
	"github.com/defn/cloud/generated/kreuzwerker/docker"

	"github.com/aws/constructs-go/constructs/v10"
	j "github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewDockerStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	docker.NewDockerProvider(stack, j.String("provider"), &docker.DockerProviderConfig{})

	dockerImage := docker.NewImage(stack, j.String("nginxImage"), &docker.ImageConfig{
		Name:        j.String("nginx:latest"),
		KeepLocally: j.Bool(false),
	})

	docker.NewContainer(stack, j.String("nginxContainer"), &docker.ContainerConfig{
		Image: dockerImage.Latest(),
		Name:  j.String("tutorial"),
		Ports: &[]*docker.ContainerPorts{{
			Internal: j.Number(80), External: j.Number(8000),
		}},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewDockerStack(app, "DockerStack")

	app.Synth()
}
