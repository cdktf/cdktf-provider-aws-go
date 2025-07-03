// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serverlessapplicationrepositorycloudformationstack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerlessapplicationrepositoryCloudformationStackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#application_id ServerlessapplicationrepositoryCloudformationStack#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#capabilities ServerlessapplicationrepositoryCloudformationStack#capabilities}.
	Capabilities *[]*string `field:"required" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#name ServerlessapplicationrepositoryCloudformationStack#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#id ServerlessapplicationrepositoryCloudformationStack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#parameters ServerlessapplicationrepositoryCloudformationStack#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#region ServerlessapplicationrepositoryCloudformationStack#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#semantic_version ServerlessapplicationrepositoryCloudformationStack#semantic_version}.
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#tags ServerlessapplicationrepositoryCloudformationStack#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#tags_all ServerlessapplicationrepositoryCloudformationStack#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#timeouts ServerlessapplicationrepositoryCloudformationStack#timeouts}
	Timeouts *ServerlessapplicationrepositoryCloudformationStackTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

