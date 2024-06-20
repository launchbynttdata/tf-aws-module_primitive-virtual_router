package common

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

func TestDoesAppmeshRouterExist(t *testing.T, ctx types.TestContext) {
	appmeshClient := appmesh.NewFromConfig(GetAWSConfig(t))
	routerName := terraform.Output(t, ctx.TerratestTerraformOptions(), "name")
	meshName := terraform.Output(t, ctx.TerratestTerraformOptions(), "mesh_name")

	output, err := appmeshClient.DescribeVirtualRouter(context.TODO(), &appmesh.DescribeVirtualRouterInput{
		MeshName:          &meshName,
		VirtualRouterName: &routerName,
	})
	if err != nil {
		t.Fatalf("unable to describe virtual router, %v", err)
	}

	t.Run("DoesAppmeshRouterExist", func(t *testing.T) {
		require.Equal(t, routerName, *output.VirtualRouter.VirtualRouterName, "Expected router name to be %s, but got %s", routerName, *output.VirtualRouter.VirtualRouterName)
		require.Equal(t, "ACTIVE", string(output.VirtualRouter.Status.Status), "Expected router status to be ACTIVE, but got %s", string(output.VirtualRouter.Status.Status))
	})

	t.Run("DoListenersExist", func(t *testing.T) {
		require.Equal(t, 2, len(output.VirtualRouter.Spec.Listeners), "Expected 2 listeners, but got %d", len(output.VirtualRouter.Spec.Listeners))
	})

}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
