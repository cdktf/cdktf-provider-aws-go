package kms

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS Key Management Service.
type DataAwsKmsPublicKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/kms_public_key#key_id DataAwsKmsPublicKey#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/kms_public_key#grant_tokens DataAwsKmsPublicKey#grant_tokens}.
	GrantTokens *[]*string `field:"optional" json:"grantTokens" yaml:"grantTokens"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/kms_public_key#id DataAwsKmsPublicKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
