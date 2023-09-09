// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrstudio

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrStudioConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#auth_mode EmrStudio#auth_mode}.
	AuthMode *string `field:"required" json:"authMode" yaml:"authMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#default_s3_location EmrStudio#default_s3_location}.
	DefaultS3Location *string `field:"required" json:"defaultS3Location" yaml:"defaultS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#engine_security_group_id EmrStudio#engine_security_group_id}.
	EngineSecurityGroupId *string `field:"required" json:"engineSecurityGroupId" yaml:"engineSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#name EmrStudio#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#service_role EmrStudio#service_role}.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#subnet_ids EmrStudio#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#vpc_id EmrStudio#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#workspace_security_group_id EmrStudio#workspace_security_group_id}.
	WorkspaceSecurityGroupId *string `field:"required" json:"workspaceSecurityGroupId" yaml:"workspaceSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#description EmrStudio#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#id EmrStudio#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#idp_auth_url EmrStudio#idp_auth_url}.
	IdpAuthUrl *string `field:"optional" json:"idpAuthUrl" yaml:"idpAuthUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#idp_relay_state_parameter_name EmrStudio#idp_relay_state_parameter_name}.
	IdpRelayStateParameterName *string `field:"optional" json:"idpRelayStateParameterName" yaml:"idpRelayStateParameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#tags EmrStudio#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#tags_all EmrStudio#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/emr_studio#user_role EmrStudio#user_role}.
	UserRole *string `field:"optional" json:"userRole" yaml:"userRole"`
}

