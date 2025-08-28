// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamWorkerAccessConfigurationS3Presign struct {
	// iam_policy_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/sagemaker_workteam#iam_policy_constraints SagemakerWorkteam#iam_policy_constraints}
	IamPolicyConstraints *SagemakerWorkteamWorkerAccessConfigurationS3PresignIamPolicyConstraints `field:"optional" json:"iamPolicyConstraints" yaml:"iamPolicyConstraints"`
}

