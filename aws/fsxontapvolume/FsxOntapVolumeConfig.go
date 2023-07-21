package fsxontapvolume

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FsxOntapVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#name FsxOntapVolume#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#size_in_megabytes FsxOntapVolume#size_in_megabytes}.
	SizeInMegabytes *float64 `field:"required" json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#storage_virtual_machine_id FsxOntapVolume#storage_virtual_machine_id}.
	StorageVirtualMachineId *string `field:"required" json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#id FsxOntapVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#junction_path FsxOntapVolume#junction_path}.
	JunctionPath *string `field:"optional" json:"junctionPath" yaml:"junctionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#ontap_volume_type FsxOntapVolume#ontap_volume_type}.
	OntapVolumeType *string `field:"optional" json:"ontapVolumeType" yaml:"ontapVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#security_style FsxOntapVolume#security_style}.
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#skip_final_backup FsxOntapVolume#skip_final_backup}.
	SkipFinalBackup interface{} `field:"optional" json:"skipFinalBackup" yaml:"skipFinalBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#storage_efficiency_enabled FsxOntapVolume#storage_efficiency_enabled}.
	StorageEfficiencyEnabled interface{} `field:"optional" json:"storageEfficiencyEnabled" yaml:"storageEfficiencyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#tags FsxOntapVolume#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#tags_all FsxOntapVolume#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tiering_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#tiering_policy FsxOntapVolume#tiering_policy}
	TieringPolicy *FsxOntapVolumeTieringPolicy `field:"optional" json:"tieringPolicy" yaml:"tieringPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#timeouts FsxOntapVolume#timeouts}
	Timeouts *FsxOntapVolumeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/fsx_ontap_volume#volume_type FsxOntapVolume#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

