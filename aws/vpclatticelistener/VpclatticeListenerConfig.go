// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticelistener

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpclatticeListenerConfig struct {
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
	// default_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#default_action VpclatticeListener#default_action}
	DefaultAction *VpclatticeListenerDefaultAction `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#name VpclatticeListener#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#protocol VpclatticeListener#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#id VpclatticeListener#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#port VpclatticeListener#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#region VpclatticeListener#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#service_arn VpclatticeListener#service_arn}.
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#service_identifier VpclatticeListener#service_identifier}.
	ServiceIdentifier *string `field:"optional" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#tags VpclatticeListener#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#tags_all VpclatticeListener#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/vpclattice_listener#timeouts VpclatticeListener#timeouts}
	Timeouts *VpclatticeListenerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

