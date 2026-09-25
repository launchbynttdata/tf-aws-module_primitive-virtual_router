package common

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	appmeshtypes "github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

type routerVerification struct {
	client     *appmesh.Client
	routerArn  string
	routerName string
	meshName   string
}

// TestComposableVirtualRouter runs read-only assertions first, then performs a
// small mutating operation (temporary tag add/remove) to prove write behavior.
func TestComposableVirtualRouter(t *testing.T, ctx types.TestContext) {
	verification := verifyRouterReadOnly(t, ctx)
	runRouterTagWriteProbe(t, verification.client, verification.routerArn)
}

// TestComposableVirtualRouterReadOnly validates the deployed virtual router via
// read-only SDK calls only. It shares verifyRouterReadOnly with the functional
// test but never invokes a mutating operation.
func TestComposableVirtualRouterReadOnly(t *testing.T, ctx types.TestContext) {
	verifyRouterReadOnly(t, ctx)
}

func verifyRouterReadOnly(t *testing.T, ctx types.TestContext) routerVerification {
	t.Helper()

	appmeshClient := appmesh.NewFromConfig(GetAWSConfig(t))
	routerName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "name")
	meshName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "mesh_name")

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

	return routerVerification{
		client:     appmeshClient,
		routerArn:  *output.VirtualRouter.Metadata.Arn,
		routerName: routerName,
		meshName:   meshName,
	}
}

// runRouterTagWriteProbe proves the module's write path by adding a temporary
// tag to the virtual router, verifying it took effect, then removing it. It
// must only be called from the functional (non-readonly) test path.
func runRouterTagWriteProbe(t *testing.T, client *appmesh.Client, routerArn string) {
	t.Run("CanTagAndUntagRouter", func(t *testing.T) {
		const probeKey = "lcaf-readonly-probe"
		const probeValue = "terratest"

		_, err := client.TagResource(context.TODO(), &appmesh.TagResourceInput{
			ResourceArn: &routerArn,
			Tags: []appmeshtypes.TagRef{
				{Key: aws.String(probeKey), Value: aws.String(probeValue)},
			},
		})
		require.NoErrorf(t, err, "unable to tag virtual router, %v", err)

		defer func() {
			_, err := client.UntagResource(context.TODO(), &appmesh.UntagResourceInput{
				ResourceArn: &routerArn,
				TagKeys:     []string{probeKey},
			})
			require.NoErrorf(t, err, "unable to untag virtual router, %v", err)
		}()

		tagsOutput, err := client.ListTagsForResource(context.TODO(), &appmesh.ListTagsForResourceInput{
			ResourceArn: &routerArn,
		})
		require.NoErrorf(t, err, "unable to list tags for virtual router, %v", err)

		var found bool
		for _, tag := range tagsOutput.Tags {
			if aws.ToString(tag.Key) == probeKey {
				require.Equal(t, probeValue, aws.ToString(tag.Value), "Expected probe tag value to be %s, but got %s", probeValue, aws.ToString(tag.Value))
				found = true
				break
			}
		}
		require.True(t, found, "Expected probe tag %s to be present after TagResource", probeKey)
	})
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
