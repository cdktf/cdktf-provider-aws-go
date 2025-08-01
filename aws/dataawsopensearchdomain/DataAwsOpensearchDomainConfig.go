// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsopensearchdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsOpensearchDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/opensearch_domain#domain_name DataAwsOpensearchDomain#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/opensearch_domain#id DataAwsOpensearchDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// off_peak_window_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/opensearch_domain#off_peak_window_options DataAwsOpensearchDomain#off_peak_window_options}
	OffPeakWindowOptions *DataAwsOpensearchDomainOffPeakWindowOptions `field:"optional" json:"offPeakWindowOptions" yaml:"offPeakWindowOptions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/opensearch_domain#region DataAwsOpensearchDomain#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/opensearch_domain#tags DataAwsOpensearchDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

