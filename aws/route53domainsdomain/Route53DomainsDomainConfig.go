// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53DomainsDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#domain_name Route53DomainsDomain#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// admin_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#admin_contact Route53DomainsDomain#admin_contact}
	AdminContact interface{} `field:"optional" json:"adminContact" yaml:"adminContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#admin_privacy Route53DomainsDomain#admin_privacy}.
	AdminPrivacy interface{} `field:"optional" json:"adminPrivacy" yaml:"adminPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#auto_renew Route53DomainsDomain#auto_renew}.
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#billing_contact Route53DomainsDomain#billing_contact}.
	BillingContact interface{} `field:"optional" json:"billingContact" yaml:"billingContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#billing_privacy Route53DomainsDomain#billing_privacy}.
	BillingPrivacy interface{} `field:"optional" json:"billingPrivacy" yaml:"billingPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#duration_in_years Route53DomainsDomain#duration_in_years}.
	DurationInYears *float64 `field:"optional" json:"durationInYears" yaml:"durationInYears"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#name_server Route53DomainsDomain#name_server}.
	NameServer interface{} `field:"optional" json:"nameServer" yaml:"nameServer"`
	// registrant_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#registrant_contact Route53DomainsDomain#registrant_contact}
	RegistrantContact interface{} `field:"optional" json:"registrantContact" yaml:"registrantContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#registrant_privacy Route53DomainsDomain#registrant_privacy}.
	RegistrantPrivacy interface{} `field:"optional" json:"registrantPrivacy" yaml:"registrantPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#tags Route53DomainsDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tech_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#tech_contact Route53DomainsDomain#tech_contact}
	TechContact interface{} `field:"optional" json:"techContact" yaml:"techContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#tech_privacy Route53DomainsDomain#tech_privacy}.
	TechPrivacy interface{} `field:"optional" json:"techPrivacy" yaml:"techPrivacy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#timeouts Route53DomainsDomain#timeouts}
	Timeouts *Route53DomainsDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/route53domains_domain#transfer_lock Route53DomainsDomain#transfer_lock}.
	TransferLock interface{} `field:"optional" json:"transferLock" yaml:"transferLock"`
}

