// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#domain_identifier DatazoneEnvironment#domain_identifier}.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#name DatazoneEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#profile_identifier DatazoneEnvironment#profile_identifier}.
	ProfileIdentifier *string `field:"required" json:"profileIdentifier" yaml:"profileIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#project_identifier DatazoneEnvironment#project_identifier}.
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#account_identifier DatazoneEnvironment#account_identifier}.
	AccountIdentifier *string `field:"optional" json:"accountIdentifier" yaml:"accountIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#account_region DatazoneEnvironment#account_region}.
	AccountRegion *string `field:"optional" json:"accountRegion" yaml:"accountRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#blueprint_identifier DatazoneEnvironment#blueprint_identifier}.
	BlueprintIdentifier *string `field:"optional" json:"blueprintIdentifier" yaml:"blueprintIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#description DatazoneEnvironment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#glossary_terms DatazoneEnvironment#glossary_terms}.
	GlossaryTerms *[]*string `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#region DatazoneEnvironment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#timeouts DatazoneEnvironment#timeouts}
	Timeouts *DatazoneEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/datazone_environment#user_parameters DatazoneEnvironment#user_parameters}
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

