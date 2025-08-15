// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangerevisionassets

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataexchangeRevisionAssetsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#data_set_id DataexchangeRevisionAssets#data_set_id}.
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
	// asset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#asset DataexchangeRevisionAssets#asset}
	Asset interface{} `field:"optional" json:"asset" yaml:"asset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#comment DataexchangeRevisionAssets#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#finalized DataexchangeRevisionAssets#finalized}.
	Finalized interface{} `field:"optional" json:"finalized" yaml:"finalized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#force_destroy DataexchangeRevisionAssets#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#region DataexchangeRevisionAssets#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#tags DataexchangeRevisionAssets#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/dataexchange_revision_assets#timeouts DataexchangeRevisionAssets#timeouts}
	Timeouts *DataexchangeRevisionAssetsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

