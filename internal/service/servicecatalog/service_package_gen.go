// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package servicecatalog

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	endpoints_sdkv1 "github.com/aws/aws-sdk-go/aws/endpoints"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	servicecatalog_sdkv1 "github.com/aws/aws-sdk-go/service/servicecatalog"
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
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
		},
		{
			Factory:  DataSourceLaunchPaths,
			TypeName: "aws_servicecatalog_launch_paths",
		},
		{
			Factory:  DataSourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  DataSourcePortfolioConstraints,
			TypeName: "aws_servicecatalog_portfolio_constraints",
		},
		{
			Factory:  DataSourceProduct,
			TypeName: "aws_servicecatalog_product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  DataSourceProvisioningArtifacts,
			TypeName: "aws_servicecatalog_provisioning_artifacts",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBudgetResourceAssociation,
			TypeName: "aws_servicecatalog_budget_resource_association",
		},
		{
			Factory:  ResourceConstraint,
			TypeName: "aws_servicecatalog_constraint",
		},
		{
			Factory:  ResourceOrganizationsAccess,
			TypeName: "aws_servicecatalog_organizations_access",
		},
		{
			Factory:  ResourcePortfolio,
			TypeName: "aws_servicecatalog_portfolio",
			Name:     "Portfolio",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourcePortfolioShare,
			TypeName: "aws_servicecatalog_portfolio_share",
		},
		{
			Factory:  ResourcePrincipalPortfolioAssociation,
			TypeName: "aws_servicecatalog_principal_portfolio_association",
		},
		{
			Factory:  ResourceProduct,
			TypeName: "aws_servicecatalog_product",
			Name:     "Product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceProductPortfolioAssociation,
			TypeName: "aws_servicecatalog_product_portfolio_association",
		},
		{
			Factory:  ResourceProvisionedProduct,
			TypeName: "aws_servicecatalog_provisioned_product",
			Name:     "Provisioned Product",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceProvisioningArtifact,
			TypeName: "aws_servicecatalog_provisioning_artifact",
		},
		{
			Factory:  ResourceServiceAction,
			TypeName: "aws_servicecatalog_service_action",
		},
		{
			Factory:  ResourceTagOption,
			TypeName: "aws_servicecatalog_tag_option",
		},
		{
			Factory:  ResourceTagOptionResourceAssociation,
			TypeName: "aws_servicecatalog_tag_option_resource_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceCatalog
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*servicecatalog_sdkv1.ServiceCatalog, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	cfg := aws_sdkv1.Config{}

	if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
		tflog.Debug(ctx, "setting endpoint", map[string]any{
			"tf_aws.endpoint": endpoint,
		})
		cfg.Endpoint = aws_sdkv1.String(endpoint)

		if sess.Config.UseFIPSEndpoint == endpoints_sdkv1.FIPSEndpointStateEnabled {
			tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
			cfg.UseFIPSEndpoint = endpoints_sdkv1.FIPSEndpointStateDisabled
		}
	}

	return servicecatalog_sdkv1.New(sess.Copy(&cfg)), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
