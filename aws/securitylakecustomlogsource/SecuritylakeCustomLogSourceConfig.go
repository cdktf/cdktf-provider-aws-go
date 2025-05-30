// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakecustomlogsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecuritylakeCustomLogSourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securitylake_custom_log_source#source_name SecuritylakeCustomLogSource#source_name}.
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securitylake_custom_log_source#configuration SecuritylakeCustomLogSource#configuration}
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securitylake_custom_log_source#event_classes SecuritylakeCustomLogSource#event_classes}.
	EventClasses *[]*string `field:"optional" json:"eventClasses" yaml:"eventClasses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/securitylake_custom_log_source#source_version SecuritylakeCustomLogSource#source_version}.
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

