// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryserviceradiussettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DirectoryServiceRadiusSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#authentication_protocol DirectoryServiceRadiusSettings#authentication_protocol}.
	AuthenticationProtocol *string `field:"required" json:"authenticationProtocol" yaml:"authenticationProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#directory_id DirectoryServiceRadiusSettings#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#display_label DirectoryServiceRadiusSettings#display_label}.
	DisplayLabel *string `field:"required" json:"displayLabel" yaml:"displayLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#radius_port DirectoryServiceRadiusSettings#radius_port}.
	RadiusPort *float64 `field:"required" json:"radiusPort" yaml:"radiusPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#radius_retries DirectoryServiceRadiusSettings#radius_retries}.
	RadiusRetries *float64 `field:"required" json:"radiusRetries" yaml:"radiusRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#radius_servers DirectoryServiceRadiusSettings#radius_servers}.
	RadiusServers *[]*string `field:"required" json:"radiusServers" yaml:"radiusServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#radius_timeout DirectoryServiceRadiusSettings#radius_timeout}.
	RadiusTimeout *float64 `field:"required" json:"radiusTimeout" yaml:"radiusTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#shared_secret DirectoryServiceRadiusSettings#shared_secret}.
	SharedSecret *string `field:"required" json:"sharedSecret" yaml:"sharedSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#id DirectoryServiceRadiusSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#timeouts DirectoryServiceRadiusSettings#timeouts}
	Timeouts *DirectoryServiceRadiusSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/directory_service_radius_settings#use_same_username DirectoryServiceRadiusSettings#use_same_username}.
	UseSameUsername interface{} `field:"optional" json:"useSameUsername" yaml:"useSameUsername"`
}

