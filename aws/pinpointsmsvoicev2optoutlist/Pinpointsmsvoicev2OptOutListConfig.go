// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2optoutlist

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Pinpointsmsvoicev2OptOutListConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/pinpointsmsvoicev2_opt_out_list#name Pinpointsmsvoicev2OptOutList#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/pinpointsmsvoicev2_opt_out_list#tags Pinpointsmsvoicev2OptOutList#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

