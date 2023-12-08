// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectvocabulary


type ConnectVocabularyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/connect_vocabulary#create ConnectVocabulary#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/connect_vocabulary#delete ConnectVocabulary#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

