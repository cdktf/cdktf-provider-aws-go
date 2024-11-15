// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53profilesresourceassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ProfilesResourceAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/route53profiles_resource_association#name Route53ProfilesResourceAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/route53profiles_resource_association#profile_id Route53ProfilesResourceAssociation#profile_id}.
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/route53profiles_resource_association#resource_arn Route53ProfilesResourceAssociation#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/route53profiles_resource_association#resource_properties Route53ProfilesResourceAssociation#resource_properties}.
	ResourceProperties *string `field:"optional" json:"resourceProperties" yaml:"resourceProperties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/route53profiles_resource_association#timeouts Route53ProfilesResourceAssociation#timeouts}
	Timeouts *Route53ProfilesResourceAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

