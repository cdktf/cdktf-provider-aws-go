// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type MemorydbUserAuthenticationMode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/memorydb_user#passwords MemorydbUser#passwords}.
	Passwords *[]*string `field:"required" json:"passwords" yaml:"passwords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/memorydb_user#type MemorydbUser#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

