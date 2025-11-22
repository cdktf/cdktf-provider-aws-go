// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightiampolicyassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightIamPolicyAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/quicksight_iam_policy_assignment#assignment_name QuicksightIamPolicyAssignment#assignment_name}.
	AssignmentName *string `field:"required" json:"assignmentName" yaml:"assignmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/quicksight_iam_policy_assignment#assignment_status QuicksightIamPolicyAssignment#assignment_status}.
	AssignmentStatus *string `field:"required" json:"assignmentStatus" yaml:"assignmentStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/quicksight_iam_policy_assignment#aws_account_id QuicksightIamPolicyAssignment#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// identities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/quicksight_iam_policy_assignment#identities QuicksightIamPolicyAssignment#identities}
	Identities interface{} `field:"optional" json:"identities" yaml:"identities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/quicksight_iam_policy_assignment#namespace QuicksightIamPolicyAssignment#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/quicksight_iam_policy_assignment#policy_arn QuicksightIamPolicyAssignment#policy_arn}.
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/quicksight_iam_policy_assignment#region QuicksightIamPolicyAssignment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

