// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package cloudsearch

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	cloudsearch_sdkv2 "github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDomain,
			TypeName: "aws_cloudsearch_domain",
			Name:     "Domain",
		},
		{
			Factory:  resourceDomainServiceAccessPolicy,
			TypeName: "aws_cloudsearch_domain_service_access_policy",
			Name:     "Domain Service Access Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CloudSearch
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*cloudsearch_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return cloudsearch_sdkv2.NewFromConfig(cfg, func(o *cloudsearch_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			tflog.Debug(ctx, "setting endpoint", map[string]any{
				"tf_aws.endpoint": endpoint,
			})
			o.BaseEndpoint = aws_sdkv2.String(endpoint)

			if o.EndpointOptions.UseFIPSEndpoint == aws_sdkv2.FIPSEndpointStateEnabled {
				tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
				o.EndpointOptions.UseFIPSEndpoint = aws_sdkv2.FIPSEndpointStateDisabled
			}
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
