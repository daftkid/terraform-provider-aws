// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	cloudformation_sdkv2 "github.com/aws/aws-sdk-go-v2/service/cloudformation"
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
			Factory:  DataSourceExport,
			TypeName: "aws_cloudformation_export",
		},
		{
			Factory:  DataSourceStack,
			TypeName: "aws_cloudformation_stack",
		},
		{
			Factory:  DataSourceType,
			TypeName: "aws_cloudformation_type",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceStack,
			TypeName: "aws_cloudformation_stack",
			Name:     "Stack",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceStackSet,
			TypeName: "aws_cloudformation_stack_set",
			Name:     "Stack Set",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceStackSetInstance,
			TypeName: "aws_cloudformation_stack_set_instance",
		},
		{
			Factory:  resourceType,
			TypeName: "aws_cloudformation_type",
			Name:     "Type",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudFormation
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*cloudformation_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return cloudformation_sdkv2.NewFromConfig(cfg, func(o *cloudformation_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
