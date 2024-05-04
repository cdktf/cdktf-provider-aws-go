// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessresourceset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoveryreadinessResourceSetConfig struct {
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
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/route53recoveryreadiness_resource_set#resources Route53RecoveryreadinessResourceSet#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/route53recoveryreadiness_resource_set#resource_set_name Route53RecoveryreadinessResourceSet#resource_set_name}.
	ResourceSetName *string `field:"required" json:"resourceSetName" yaml:"resourceSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/route53recoveryreadiness_resource_set#resource_set_type Route53RecoveryreadinessResourceSet#resource_set_type}.
	ResourceSetType *string `field:"required" json:"resourceSetType" yaml:"resourceSetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/route53recoveryreadiness_resource_set#id Route53RecoveryreadinessResourceSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/route53recoveryreadiness_resource_set#tags Route53RecoveryreadinessResourceSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/route53recoveryreadiness_resource_set#tags_all Route53RecoveryreadinessResourceSet#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/route53recoveryreadiness_resource_set#timeouts Route53RecoveryreadinessResourceSet#timeouts}
	Timeouts *Route53RecoveryreadinessResourceSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

