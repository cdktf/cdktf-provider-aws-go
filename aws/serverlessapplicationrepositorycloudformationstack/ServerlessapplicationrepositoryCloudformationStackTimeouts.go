// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serverlessapplicationrepositorycloudformationstack


type ServerlessapplicationrepositoryCloudformationStackTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/serverlessapplicationrepository_cloudformation_stack#create ServerlessapplicationrepositoryCloudformationStack#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/serverlessapplicationrepository_cloudformation_stack#delete ServerlessapplicationrepositoryCloudformationStack#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/serverlessapplicationrepository_cloudformation_stack#update ServerlessapplicationrepositoryCloudformationStack#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

