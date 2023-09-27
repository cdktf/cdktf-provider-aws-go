// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipamresourcediscoveryassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcIpamResourceDiscoveryAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/vpc_ipam_resource_discovery_association#ipam_id VpcIpamResourceDiscoveryAssociation#ipam_id}.
	IpamId *string `field:"required" json:"ipamId" yaml:"ipamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/vpc_ipam_resource_discovery_association#ipam_resource_discovery_id VpcIpamResourceDiscoveryAssociation#ipam_resource_discovery_id}.
	IpamResourceDiscoveryId *string `field:"required" json:"ipamResourceDiscoveryId" yaml:"ipamResourceDiscoveryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/vpc_ipam_resource_discovery_association#id VpcIpamResourceDiscoveryAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/vpc_ipam_resource_discovery_association#tags VpcIpamResourceDiscoveryAssociation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/vpc_ipam_resource_discovery_association#tags_all VpcIpamResourceDiscoveryAssociation#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/vpc_ipam_resource_discovery_association#timeouts VpcIpamResourceDiscoveryAssociation#timeouts}
	Timeouts *VpcIpamResourceDiscoveryAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

