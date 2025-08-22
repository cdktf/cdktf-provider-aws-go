// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerWorkteamConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#description SagemakerWorkteam#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// member_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#member_definition SagemakerWorkteam#member_definition}
	MemberDefinition interface{} `field:"required" json:"memberDefinition" yaml:"memberDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#workteam_name SagemakerWorkteam#workteam_name}.
	WorkteamName *string `field:"required" json:"workteamName" yaml:"workteamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#id SagemakerWorkteam#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// notification_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#notification_configuration SagemakerWorkteam#notification_configuration}
	NotificationConfiguration *SagemakerWorkteamNotificationConfiguration `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#region SagemakerWorkteam#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#tags SagemakerWorkteam#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#tags_all SagemakerWorkteam#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// worker_access_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#worker_access_configuration SagemakerWorkteam#worker_access_configuration}
	WorkerAccessConfiguration *SagemakerWorkteamWorkerAccessConfiguration `field:"optional" json:"workerAccessConfiguration" yaml:"workerAccessConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_workteam#workforce_name SagemakerWorkteam#workforce_name}.
	WorkforceName *string `field:"optional" json:"workforceName" yaml:"workforceName"`
}

