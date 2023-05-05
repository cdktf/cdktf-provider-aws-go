package datasynclocations3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationS3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#s3_bucket_arn DatasyncLocationS3#s3_bucket_arn}.
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#s3_config DatasyncLocationS3#s3_config}
	S3Config *DatasyncLocationS3S3Config `field:"required" json:"s3Config" yaml:"s3Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#subdirectory DatasyncLocationS3#subdirectory}.
	Subdirectory *string `field:"required" json:"subdirectory" yaml:"subdirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#agent_arns DatasyncLocationS3#agent_arns}.
	AgentArns *[]*string `field:"optional" json:"agentArns" yaml:"agentArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#id DatasyncLocationS3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#s3_storage_class DatasyncLocationS3#s3_storage_class}.
	S3StorageClass *string `field:"optional" json:"s3StorageClass" yaml:"s3StorageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#tags DatasyncLocationS3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/datasync_location_s3#tags_all DatasyncLocationS3#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

