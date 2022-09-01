package gamelift

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS GameLift.
type GameliftScriptConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_script#name GameliftScript#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_script#id GameliftScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// storage_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_script#storage_location GameliftScript#storage_location}
	StorageLocation *GameliftScriptStorageLocation `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_script#tags GameliftScript#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_script#tags_all GameliftScript#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_script#version GameliftScript#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/gamelift_script#zip_file GameliftScript#zip_file}.
	ZipFile *string `field:"optional" json:"zipFile" yaml:"zipFile"`
}

