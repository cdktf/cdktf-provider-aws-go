// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowConnectorProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appflow_connector_profile#connection_mode AppflowConnectorProfile#connection_mode}.
	ConnectionMode *string `field:"required" json:"connectionMode" yaml:"connectionMode"`
	// connector_profile_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appflow_connector_profile#connector_profile_config AppflowConnectorProfile#connector_profile_config}
	ConnectorProfileConfig *AppflowConnectorProfileConnectorProfileConfig `field:"required" json:"connectorProfileConfig" yaml:"connectorProfileConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appflow_connector_profile#connector_type AppflowConnectorProfile#connector_type}.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appflow_connector_profile#name AppflowConnectorProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appflow_connector_profile#connector_label AppflowConnectorProfile#connector_label}.
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appflow_connector_profile#id AppflowConnectorProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/appflow_connector_profile#kms_arn AppflowConnectorProfile#kms_arn}.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}

