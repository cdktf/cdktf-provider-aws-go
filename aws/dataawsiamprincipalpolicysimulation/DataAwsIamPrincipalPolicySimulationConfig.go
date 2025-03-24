// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiamprincipalpolicysimulation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsIamPrincipalPolicySimulationConfig struct {
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
	// One or more names of actions, like "iam:CreateUser", that should be included in the simulation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#action_names DataAwsIamPrincipalPolicySimulation#action_names}
	ActionNames *[]*string `field:"required" json:"actionNames" yaml:"actionNames"`
	// ARN of the principal (e.g. user, role) whose existing configured access policies will be used as the basis for the simulation. If you specify a role ARN here, you can also set caller_arn to simulate a particular user acting with the given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#policy_source_arn DataAwsIamPrincipalPolicySimulation#policy_source_arn}
	PolicySourceArn *string `field:"required" json:"policySourceArn" yaml:"policySourceArn"`
	// Additional principal-based policies to use in the simulation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#additional_policies_json DataAwsIamPrincipalPolicySimulation#additional_policies_json}
	AdditionalPoliciesJson *[]*string `field:"optional" json:"additionalPoliciesJson" yaml:"additionalPoliciesJson"`
	// ARN of a user to use as the caller of the simulated requests.
	//
	// If not specified, defaults to the principal specified in policy_source_arn, if it is a user ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#caller_arn DataAwsIamPrincipalPolicySimulation#caller_arn}
	CallerArn *string `field:"optional" json:"callerArn" yaml:"callerArn"`
	// context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#context DataAwsIamPrincipalPolicySimulation#context}
	Context interface{} `field:"optional" json:"context" yaml:"context"`
	// Additional permission boundary policies to use in the simulation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#permissions_boundary_policies_json DataAwsIamPrincipalPolicySimulation#permissions_boundary_policies_json}
	PermissionsBoundaryPoliciesJson *[]*string `field:"optional" json:"permissionsBoundaryPoliciesJson" yaml:"permissionsBoundaryPoliciesJson"`
	// ARNs of specific resources to use as the targets of the specified actions during simulation.
	//
	// If not specified, the simulator assumes "*" which represents general access across all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#resource_arns DataAwsIamPrincipalPolicySimulation#resource_arns}
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Specifies the type of simulation to run.
	//
	// Some API operations need a particular resource handling option in order to produce a correct reesult.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#resource_handling_option DataAwsIamPrincipalPolicySimulation#resource_handling_option}
	ResourceHandlingOption *string `field:"optional" json:"resourceHandlingOption" yaml:"resourceHandlingOption"`
	// An AWS account ID to use as the simulated owner for any resource whose ARN does not include a specific owner account ID.
	//
	// Defaults to the account given as part of caller_arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#resource_owner_account_id DataAwsIamPrincipalPolicySimulation#resource_owner_account_id}
	ResourceOwnerAccountId *string `field:"optional" json:"resourceOwnerAccountId" yaml:"resourceOwnerAccountId"`
	// A resource policy to associate with all of the target resources for simulation purposes.
	//
	// The policy simulator does not automatically retrieve resource-level policies, so if a resource policy is crucial to your test then you must specify here the same policy document associated with your target resource(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/data-sources/iam_principal_policy_simulation#resource_policy_json DataAwsIamPrincipalPolicySimulation#resource_policy_json}
	ResourcePolicyJson *string `field:"optional" json:"resourcePolicyJson" yaml:"resourcePolicyJson"`
}

