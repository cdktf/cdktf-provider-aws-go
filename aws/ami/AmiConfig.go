package ami

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AmiConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#name Ami#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#architecture Ami#architecture}.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#boot_mode Ami#boot_mode}.
	BootMode *string `field:"optional" json:"bootMode" yaml:"bootMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#deprecation_time Ami#deprecation_time}.
	DeprecationTime *string `field:"optional" json:"deprecationTime" yaml:"deprecationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#description Ami#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#ebs_block_device Ami#ebs_block_device}
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#ena_support Ami#ena_support}.
	EnaSupport interface{} `field:"optional" json:"enaSupport" yaml:"enaSupport"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#ephemeral_block_device Ami#ephemeral_block_device}
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#id Ami#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#image_location Ami#image_location}.
	ImageLocation *string `field:"optional" json:"imageLocation" yaml:"imageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#imds_support Ami#imds_support}.
	ImdsSupport *string `field:"optional" json:"imdsSupport" yaml:"imdsSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#kernel_id Ami#kernel_id}.
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#ramdisk_id Ami#ramdisk_id}.
	RamdiskId *string `field:"optional" json:"ramdiskId" yaml:"ramdiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#root_device_name Ami#root_device_name}.
	RootDeviceName *string `field:"optional" json:"rootDeviceName" yaml:"rootDeviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#sriov_net_support Ami#sriov_net_support}.
	SriovNetSupport *string `field:"optional" json:"sriovNetSupport" yaml:"sriovNetSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#tags Ami#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#tags_all Ami#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#timeouts Ami#timeouts}
	Timeouts *AmiTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#tpm_support Ami#tpm_support}.
	TpmSupport *string `field:"optional" json:"tpmSupport" yaml:"tpmSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/ami#virtualization_type Ami#virtualization_type}.
	VirtualizationType *string `field:"optional" json:"virtualizationType" yaml:"virtualizationType"`
}

