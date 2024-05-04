// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewayfilesystemassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StoragegatewayFileSystemAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#gateway_arn StoragegatewayFileSystemAssociation#gateway_arn}.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#location_arn StoragegatewayFileSystemAssociation#location_arn}.
	LocationArn *string `field:"required" json:"locationArn" yaml:"locationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#password StoragegatewayFileSystemAssociation#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#username StoragegatewayFileSystemAssociation#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#audit_destination_arn StoragegatewayFileSystemAssociation#audit_destination_arn}.
	AuditDestinationArn *string `field:"optional" json:"auditDestinationArn" yaml:"auditDestinationArn"`
	// cache_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#cache_attributes StoragegatewayFileSystemAssociation#cache_attributes}
	CacheAttributes *StoragegatewayFileSystemAssociationCacheAttributes `field:"optional" json:"cacheAttributes" yaml:"cacheAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#id StoragegatewayFileSystemAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#tags StoragegatewayFileSystemAssociation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/storagegateway_file_system_association#tags_all StoragegatewayFileSystemAssociation#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

