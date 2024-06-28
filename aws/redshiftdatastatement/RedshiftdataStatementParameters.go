// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftdatastatement


type RedshiftdataStatementParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/redshiftdata_statement#name RedshiftdataStatement#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/redshiftdata_statement#value RedshiftdataStatement#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

