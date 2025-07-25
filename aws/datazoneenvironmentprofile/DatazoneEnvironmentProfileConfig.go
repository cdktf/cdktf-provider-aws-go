// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironmentprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneEnvironmentProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#aws_account_region DatazoneEnvironmentProfile#aws_account_region}.
	AwsAccountRegion *string `field:"required" json:"awsAccountRegion" yaml:"awsAccountRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#domain_identifier DatazoneEnvironmentProfile#domain_identifier}.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#environment_blueprint_identifier DatazoneEnvironmentProfile#environment_blueprint_identifier}.
	EnvironmentBlueprintIdentifier *string `field:"required" json:"environmentBlueprintIdentifier" yaml:"environmentBlueprintIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#name DatazoneEnvironmentProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#project_identifier DatazoneEnvironmentProfile#project_identifier}.
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#aws_account_id DatazoneEnvironmentProfile#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#description DatazoneEnvironmentProfile#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#region DatazoneEnvironmentProfile#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// user_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/datazone_environment_profile#user_parameters DatazoneEnvironmentProfile#user_parameters}
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

