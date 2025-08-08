// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudsearchdomainserviceaccesspolicy


type CloudsearchDomainServiceAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/cloudsearch_domain_service_access_policy#delete CloudsearchDomainServiceAccessPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/cloudsearch_domain_service_access_policy#update CloudsearchDomainServiceAccessPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

