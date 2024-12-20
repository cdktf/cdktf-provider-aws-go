// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamservercertificate


type IamServerCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/iam_server_certificate#delete IamServerCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

