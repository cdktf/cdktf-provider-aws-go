// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcrouteserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcRouteServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_route_server#amazon_side_asn VpcRouteServer#amazon_side_asn}.
	AmazonSideAsn *float64 `field:"required" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_route_server#persist_routes VpcRouteServer#persist_routes}.
	PersistRoutes *string `field:"optional" json:"persistRoutes" yaml:"persistRoutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_route_server#persist_routes_duration VpcRouteServer#persist_routes_duration}.
	PersistRoutesDuration *float64 `field:"optional" json:"persistRoutesDuration" yaml:"persistRoutesDuration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_route_server#region VpcRouteServer#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_route_server#sns_notifications_enabled VpcRouteServer#sns_notifications_enabled}.
	SnsNotificationsEnabled interface{} `field:"optional" json:"snsNotificationsEnabled" yaml:"snsNotificationsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_route_server#tags VpcRouteServer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/vpc_route_server#timeouts VpcRouteServer#timeouts}
	Timeouts *VpcRouteServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

