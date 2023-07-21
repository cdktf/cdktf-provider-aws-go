package dataawscloudcontrolapiresource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsCloudcontrolapiResourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/cloudcontrolapi_resource#identifier DataAwsCloudcontrolapiResource#identifier}.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/cloudcontrolapi_resource#type_name DataAwsCloudcontrolapiResource#type_name}.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/cloudcontrolapi_resource#id DataAwsCloudcontrolapiResource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/cloudcontrolapi_resource#role_arn DataAwsCloudcontrolapiResource#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/cloudcontrolapi_resource#type_version_id DataAwsCloudcontrolapiResource#type_version_id}.
	TypeVersionId *string `field:"optional" json:"typeVersionId" yaml:"typeVersionId"`
}

