// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recordsexclusive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecordsExclusiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_records_exclusive#zone_id Route53RecordsExclusive#zone_id}.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// resource_record_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_records_exclusive#resource_record_set Route53RecordsExclusive#resource_record_set}
	ResourceRecordSet interface{} `field:"optional" json:"resourceRecordSet" yaml:"resourceRecordSet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_records_exclusive#timeouts Route53RecordsExclusive#timeouts}
	Timeouts *Route53RecordsExclusiveTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

