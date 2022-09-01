package lightsail

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS Lightsail.
type LightsailInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#availability_zone LightsailInstance#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#blueprint_id LightsailInstance#blueprint_id}.
	BlueprintId *string `field:"required" json:"blueprintId" yaml:"blueprintId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#bundle_id LightsailInstance#bundle_id}.
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#name LightsailInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#id LightsailInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#key_pair_name LightsailInstance#key_pair_name}.
	KeyPairName *string `field:"optional" json:"keyPairName" yaml:"keyPairName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#tags LightsailInstance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#tags_all LightsailInstance#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lightsail_instance#user_data LightsailInstance#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

