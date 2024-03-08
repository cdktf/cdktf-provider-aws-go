// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdocdborderabledbinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDocdbOrderableDbInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/docdb_orderable_db_instance#engine DataAwsDocdbOrderableDbInstance#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/docdb_orderable_db_instance#engine_version DataAwsDocdbOrderableDbInstance#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/docdb_orderable_db_instance#id DataAwsDocdbOrderableDbInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/docdb_orderable_db_instance#instance_class DataAwsDocdbOrderableDbInstance#instance_class}.
	InstanceClass *string `field:"optional" json:"instanceClass" yaml:"instanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/docdb_orderable_db_instance#license_model DataAwsDocdbOrderableDbInstance#license_model}.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/docdb_orderable_db_instance#preferred_instance_classes DataAwsDocdbOrderableDbInstance#preferred_instance_classes}.
	PreferredInstanceClasses *[]*string `field:"optional" json:"preferredInstanceClasses" yaml:"preferredInstanceClasses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/data-sources/docdb_orderable_db_instance#vpc DataAwsDocdbOrderableDbInstance#vpc}.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

