// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener


type LbListenerDefaultActionFixedResponse struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lb_listener#content_type LbListener#content_type}.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lb_listener#message_body LbListener#message_body}.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lb_listener#status_code LbListener#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

