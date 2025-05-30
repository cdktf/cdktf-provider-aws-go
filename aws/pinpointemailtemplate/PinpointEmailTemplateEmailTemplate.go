// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointemailtemplate


type PinpointEmailTemplateEmailTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pinpoint_email_template#default_substitutions PinpointEmailTemplate#default_substitutions}.
	DefaultSubstitutions *string `field:"optional" json:"defaultSubstitutions" yaml:"defaultSubstitutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pinpoint_email_template#description PinpointEmailTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pinpoint_email_template#header PinpointEmailTemplate#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pinpoint_email_template#html_part PinpointEmailTemplate#html_part}.
	HtmlPart *string `field:"optional" json:"htmlPart" yaml:"htmlPart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pinpoint_email_template#recommender_id PinpointEmailTemplate#recommender_id}.
	RecommenderId *string `field:"optional" json:"recommenderId" yaml:"recommenderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pinpoint_email_template#subject PinpointEmailTemplate#subject}.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/pinpoint_email_template#text_part PinpointEmailTemplate#text_part}.
	TextPart *string `field:"optional" json:"textPart" yaml:"textPart"`
}

