// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionSsmParameterConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/imagebuilder_distribution_configuration#parameter_name ImagebuilderDistributionConfiguration#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/imagebuilder_distribution_configuration#ami_account_id ImagebuilderDistributionConfiguration#ami_account_id}.
	AmiAccountId *string `field:"optional" json:"amiAccountId" yaml:"amiAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/imagebuilder_distribution_configuration#data_type ImagebuilderDistributionConfiguration#data_type}.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
}

