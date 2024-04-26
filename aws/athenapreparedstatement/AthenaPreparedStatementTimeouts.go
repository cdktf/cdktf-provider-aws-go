// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package athenapreparedstatement


type AthenaPreparedStatementTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/athena_prepared_statement#create AthenaPreparedStatement#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/athena_prepared_statement#delete AthenaPreparedStatement#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/athena_prepared_statement#update AthenaPreparedStatement#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

