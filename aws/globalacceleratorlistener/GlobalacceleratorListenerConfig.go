// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorlistener

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorListenerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/globalaccelerator_listener#accelerator_arn GlobalacceleratorListener#accelerator_arn}.
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/globalaccelerator_listener#port_range GlobalacceleratorListener#port_range}
	PortRange interface{} `field:"required" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/globalaccelerator_listener#protocol GlobalacceleratorListener#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/globalaccelerator_listener#client_affinity GlobalacceleratorListener#client_affinity}.
	ClientAffinity *string `field:"optional" json:"clientAffinity" yaml:"clientAffinity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/globalaccelerator_listener#id GlobalacceleratorListener#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/globalaccelerator_listener#timeouts GlobalacceleratorListener#timeouts}
	Timeouts *GlobalacceleratorListenerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

