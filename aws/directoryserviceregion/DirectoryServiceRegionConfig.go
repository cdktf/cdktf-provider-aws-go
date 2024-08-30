// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryserviceregion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectoryServiceRegionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#directory_id DirectoryServiceRegion#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#region_name DirectoryServiceRegion#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// vpc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#vpc_settings DirectoryServiceRegion#vpc_settings}
	VpcSettings *DirectoryServiceRegionVpcSettings `field:"required" json:"vpcSettings" yaml:"vpcSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#desired_number_of_domain_controllers DirectoryServiceRegion#desired_number_of_domain_controllers}.
	DesiredNumberOfDomainControllers *float64 `field:"optional" json:"desiredNumberOfDomainControllers" yaml:"desiredNumberOfDomainControllers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#id DirectoryServiceRegion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#tags DirectoryServiceRegion#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#tags_all DirectoryServiceRegion#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/directory_service_region#timeouts DirectoryServiceRegion#timeouts}
	Timeouts *DirectoryServiceRegionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

