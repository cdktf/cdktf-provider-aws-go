package ec2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS EC2.
type DataAwsInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#get_password_data DataAwsInstance#get_password_data}.
	FetchPasswordData interface{} `field:"optional" json:"fetchPasswordData" yaml:"fetchPasswordData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#get_user_data DataAwsInstance#get_user_data}.
	FetchUserData interface{} `field:"optional" json:"fetchUserData" yaml:"fetchUserData"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#filter DataAwsInstance#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#id DataAwsInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#instance_id DataAwsInstance#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#instance_tags DataAwsInstance#instance_tags}.
	InstanceTags *map[string]*string `field:"optional" json:"instanceTags" yaml:"instanceTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#tags DataAwsInstance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/instance#timeouts DataAwsInstance#timeouts}
	Timeouts *DataAwsInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

