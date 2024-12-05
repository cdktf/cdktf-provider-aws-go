// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package glueclassifier


type GlueClassifierCsvClassifier struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#allow_single_column GlueClassifier#allow_single_column}.
	AllowSingleColumn interface{} `field:"optional" json:"allowSingleColumn" yaml:"allowSingleColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#contains_header GlueClassifier#contains_header}.
	ContainsHeader *string `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#custom_datatype_configured GlueClassifier#custom_datatype_configured}.
	CustomDatatypeConfigured interface{} `field:"optional" json:"customDatatypeConfigured" yaml:"customDatatypeConfigured"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#custom_datatypes GlueClassifier#custom_datatypes}.
	CustomDatatypes *[]*string `field:"optional" json:"customDatatypes" yaml:"customDatatypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#delimiter GlueClassifier#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#disable_value_trimming GlueClassifier#disable_value_trimming}.
	DisableValueTrimming interface{} `field:"optional" json:"disableValueTrimming" yaml:"disableValueTrimming"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#header GlueClassifier#header}.
	Header *[]*string `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#quote_symbol GlueClassifier#quote_symbol}.
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/glue_classifier#serde GlueClassifier#serde}.
	Serde *string `field:"optional" json:"serde" yaml:"serde"`
}

