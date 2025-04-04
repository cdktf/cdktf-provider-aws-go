// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanroomsMembershipConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/cleanrooms_membership#collaboration_id CleanroomsMembership#collaboration_id}.
	CollaborationId *string `field:"required" json:"collaborationId" yaml:"collaborationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/cleanrooms_membership#query_log_status CleanroomsMembership#query_log_status}.
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// default_result_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/cleanrooms_membership#default_result_configuration CleanroomsMembership#default_result_configuration}
	DefaultResultConfiguration interface{} `field:"optional" json:"defaultResultConfiguration" yaml:"defaultResultConfiguration"`
	// payment_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/cleanrooms_membership#payment_configuration CleanroomsMembership#payment_configuration}
	PaymentConfiguration interface{} `field:"optional" json:"paymentConfiguration" yaml:"paymentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/cleanrooms_membership#tags CleanroomsMembership#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

