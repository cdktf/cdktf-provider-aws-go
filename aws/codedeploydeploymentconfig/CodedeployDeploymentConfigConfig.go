// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodedeployDeploymentConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/codedeploy_deployment_config#deployment_config_name CodedeployDeploymentConfig#deployment_config_name}.
	DeploymentConfigName *string `field:"required" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/codedeploy_deployment_config#compute_platform CodedeployDeploymentConfig#compute_platform}.
	ComputePlatform *string `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/codedeploy_deployment_config#id CodedeployDeploymentConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// minimum_healthy_hosts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/codedeploy_deployment_config#minimum_healthy_hosts CodedeployDeploymentConfig#minimum_healthy_hosts}
	MinimumHealthyHosts *CodedeployDeploymentConfigMinimumHealthyHosts `field:"optional" json:"minimumHealthyHosts" yaml:"minimumHealthyHosts"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/codedeploy_deployment_config#region CodedeployDeploymentConfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// traffic_routing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/codedeploy_deployment_config#traffic_routing_config CodedeployDeploymentConfig#traffic_routing_config}
	TrafficRoutingConfig *CodedeployDeploymentConfigTrafficRoutingConfig `field:"optional" json:"trafficRoutingConfig" yaml:"trafficRoutingConfig"`
	// zonal_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/codedeploy_deployment_config#zonal_config CodedeployDeploymentConfig#zonal_config}
	ZonalConfig *CodedeployDeploymentConfigZonalConfig `field:"optional" json:"zonalConfig" yaml:"zonalConfig"`
}

