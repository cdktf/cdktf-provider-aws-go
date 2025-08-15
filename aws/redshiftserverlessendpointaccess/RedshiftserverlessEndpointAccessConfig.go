// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessendpointaccess

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftserverlessEndpointAccessConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/redshiftserverless_endpoint_access#endpoint_name RedshiftserverlessEndpointAccess#endpoint_name}.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/redshiftserverless_endpoint_access#subnet_ids RedshiftserverlessEndpointAccess#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/redshiftserverless_endpoint_access#workgroup_name RedshiftserverlessEndpointAccess#workgroup_name}.
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/redshiftserverless_endpoint_access#id RedshiftserverlessEndpointAccess#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/redshiftserverless_endpoint_access#owner_account RedshiftserverlessEndpointAccess#owner_account}.
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/redshiftserverless_endpoint_access#region RedshiftserverlessEndpointAccess#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/redshiftserverless_endpoint_access#vpc_security_group_ids RedshiftserverlessEndpointAccess#vpc_security_group_ids}.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

