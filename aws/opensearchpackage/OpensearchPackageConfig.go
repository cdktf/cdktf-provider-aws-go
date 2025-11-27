// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchpackage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchPackageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_package#package_name OpensearchPackage#package_name}.
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// package_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_package#package_source OpensearchPackage#package_source}
	PackageSource *OpensearchPackagePackageSource `field:"required" json:"packageSource" yaml:"packageSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_package#package_type OpensearchPackage#package_type}.
	PackageType *string `field:"required" json:"packageType" yaml:"packageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_package#engine_version OpensearchPackage#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_package#id OpensearchPackage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_package#package_description OpensearchPackage#package_description}.
	PackageDescription *string `field:"optional" json:"packageDescription" yaml:"packageDescription"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/opensearch_package#region OpensearchPackage#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

