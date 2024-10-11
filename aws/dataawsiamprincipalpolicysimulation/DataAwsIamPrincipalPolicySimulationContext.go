// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiamprincipalpolicysimulation


type DataAwsIamPrincipalPolicySimulationContext struct {
	// The key name of the context entry, such as "aws:CurrentTime".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/data-sources/iam_principal_policy_simulation#key DataAwsIamPrincipalPolicySimulation#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The type that the simulator should use to interpret the strings given in argument "values".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/data-sources/iam_principal_policy_simulation#type DataAwsIamPrincipalPolicySimulation#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// One or more values to assign to the context key, given as a string in a syntax appropriate for the selected value type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/data-sources/iam_principal_policy_simulation#values DataAwsIamPrincipalPolicySimulation#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

