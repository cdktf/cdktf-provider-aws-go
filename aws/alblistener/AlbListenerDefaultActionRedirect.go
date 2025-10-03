// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alblistener


type AlbListenerDefaultActionRedirect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/alb_listener#status_code AlbListener#status_code}.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/alb_listener#host AlbListener#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/alb_listener#path AlbListener#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/alb_listener#port AlbListener#port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/alb_listener#protocol AlbListener#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/alb_listener#query AlbListener#query}.
	Query *string `field:"optional" json:"query" yaml:"query"`
}

