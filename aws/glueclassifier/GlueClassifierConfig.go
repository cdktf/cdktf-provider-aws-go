// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glueclassifier

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlueClassifierConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/glue_classifier#name GlueClassifier#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// csv_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/glue_classifier#csv_classifier GlueClassifier#csv_classifier}
	CsvClassifier *GlueClassifierCsvClassifier `field:"optional" json:"csvClassifier" yaml:"csvClassifier"`
	// grok_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/glue_classifier#grok_classifier GlueClassifier#grok_classifier}
	GrokClassifier *GlueClassifierGrokClassifier `field:"optional" json:"grokClassifier" yaml:"grokClassifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/glue_classifier#id GlueClassifier#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// json_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/glue_classifier#json_classifier GlueClassifier#json_classifier}
	JsonClassifier *GlueClassifierJsonClassifier `field:"optional" json:"jsonClassifier" yaml:"jsonClassifier"`
	// xml_classifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/glue_classifier#xml_classifier GlueClassifier#xml_classifier}
	XmlClassifier *GlueClassifierXmlClassifier `field:"optional" json:"xmlClassifier" yaml:"xmlClassifier"`
}

