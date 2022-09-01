package transfer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS Transfer.
type TransferAccessConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#external_id TransferAccess#external_id}.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#server_id TransferAccess#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#home_directory TransferAccess#home_directory}.
	HomeDirectory *string `field:"optional" json:"homeDirectory" yaml:"homeDirectory"`
	// home_directory_mappings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#home_directory_mappings TransferAccess#home_directory_mappings}
	HomeDirectoryMappings interface{} `field:"optional" json:"homeDirectoryMappings" yaml:"homeDirectoryMappings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#home_directory_type TransferAccess#home_directory_type}.
	HomeDirectoryType *string `field:"optional" json:"homeDirectoryType" yaml:"homeDirectoryType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#id TransferAccess#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#policy TransferAccess#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// posix_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#posix_profile TransferAccess#posix_profile}
	PosixProfile *TransferAccessPosixProfile `field:"optional" json:"posixProfile" yaml:"posixProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#role TransferAccess#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

