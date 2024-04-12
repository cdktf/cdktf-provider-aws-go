// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitylakesubscriber

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecuritylakeSubscriberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securitylake_subscriber#access_type SecuritylakeSubscriber#access_type}.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securitylake_subscriber#source SecuritylakeSubscriber#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securitylake_subscriber#subscriber_description SecuritylakeSubscriber#subscriber_description}.
	SubscriberDescription *string `field:"optional" json:"subscriberDescription" yaml:"subscriberDescription"`
	// subscriber_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securitylake_subscriber#subscriber_identity SecuritylakeSubscriber#subscriber_identity}
	SubscriberIdentity interface{} `field:"optional" json:"subscriberIdentity" yaml:"subscriberIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securitylake_subscriber#subscriber_name SecuritylakeSubscriber#subscriber_name}.
	SubscriberName *string `field:"optional" json:"subscriberName" yaml:"subscriberName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securitylake_subscriber#tags SecuritylakeSubscriber#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/securitylake_subscriber#timeouts SecuritylakeSubscriber#timeouts}
	Timeouts *SecuritylakeSubscriberTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

