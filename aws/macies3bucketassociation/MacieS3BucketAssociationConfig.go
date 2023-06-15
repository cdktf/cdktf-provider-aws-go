package macies3bucketassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MacieS3BucketAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/macie_s3_bucket_association#bucket_name MacieS3BucketAssociation#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// classification_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/macie_s3_bucket_association#classification_type MacieS3BucketAssociation#classification_type}
	ClassificationType *MacieS3BucketAssociationClassificationType `field:"optional" json:"classificationType" yaml:"classificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/macie_s3_bucket_association#id MacieS3BucketAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/macie_s3_bucket_association#member_account_id MacieS3BucketAssociation#member_account_id}.
	MemberAccountId *string `field:"optional" json:"memberAccountId" yaml:"memberAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/macie_s3_bucket_association#prefix MacieS3BucketAssociation#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

