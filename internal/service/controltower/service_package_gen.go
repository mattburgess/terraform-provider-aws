// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package controltower

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	controltower_sdkv2 "github.com/aws/aws-sdk-go-v2/service/controltower"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceControls,
			TypeName: "aws_controltower_controls",
			Name:     "Control",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceControl,
			TypeName: "aws_controltower_control",
			Name:     "Control",
		},
		{
			Factory:  resourceLandingZone,
			TypeName: "aws_controltower_landing_zone",
			Name:     "Landing Zone",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ControlTower
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*controltower_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return controltower_sdkv2.NewFromConfig(cfg, func(o *controltower_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
