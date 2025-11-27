// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2allowedimagessettings


type Ec2AllowedImagesSettingsImageCriterion struct {
	// creation_date_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_allowed_images_settings#creation_date_condition Ec2AllowedImagesSettings#creation_date_condition}
	CreationDateCondition interface{} `field:"optional" json:"creationDateCondition" yaml:"creationDateCondition"`
	// deprecation_time_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_allowed_images_settings#deprecation_time_condition Ec2AllowedImagesSettings#deprecation_time_condition}
	DeprecationTimeCondition interface{} `field:"optional" json:"deprecationTimeCondition" yaml:"deprecationTimeCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_allowed_images_settings#image_names Ec2AllowedImagesSettings#image_names}.
	ImageNames *[]*string `field:"optional" json:"imageNames" yaml:"imageNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_allowed_images_settings#image_providers Ec2AllowedImagesSettings#image_providers}.
	ImageProviders *[]*string `field:"optional" json:"imageProviders" yaml:"imageProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ec2_allowed_images_settings#marketplace_product_codes Ec2AllowedImagesSettings#marketplace_product_codes}.
	MarketplaceProductCodes *[]*string `field:"optional" json:"marketplaceProductCodes" yaml:"marketplaceProductCodes"`
}

