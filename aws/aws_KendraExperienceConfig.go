// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KendraExperienceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#index_id KendraExperience#index_id}.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#name KendraExperience#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#role_arn KendraExperience#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#configuration KendraExperience#configuration}
	Configuration *KendraExperienceConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#description KendraExperience#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#id KendraExperience#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_experience#timeouts KendraExperience#timeouts}
	Timeouts *KendraExperienceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

