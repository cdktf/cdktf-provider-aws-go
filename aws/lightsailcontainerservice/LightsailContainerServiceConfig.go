// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainerservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LightsailContainerServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#name LightsailContainerService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#power LightsailContainerService#power}.
	Power *string `field:"required" json:"power" yaml:"power"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#scale LightsailContainerService#scale}.
	Scale *float64 `field:"required" json:"scale" yaml:"scale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#id LightsailContainerService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#is_disabled LightsailContainerService#is_disabled}.
	IsDisabled interface{} `field:"optional" json:"isDisabled" yaml:"isDisabled"`
	// private_registry_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#private_registry_access LightsailContainerService#private_registry_access}
	PrivateRegistryAccess *LightsailContainerServicePrivateRegistryAccess `field:"optional" json:"privateRegistryAccess" yaml:"privateRegistryAccess"`
	// public_domain_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#public_domain_names LightsailContainerService#public_domain_names}
	PublicDomainNames *LightsailContainerServicePublicDomainNames `field:"optional" json:"publicDomainNames" yaml:"publicDomainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#tags LightsailContainerService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#tags_all LightsailContainerService#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lightsail_container_service#timeouts LightsailContainerService#timeouts}
	Timeouts *LightsailContainerServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

