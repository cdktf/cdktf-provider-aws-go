package lightsailinstance


type LightsailInstanceAddOn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lightsail_instance#snapshot_time LightsailInstance#snapshot_time}.
	SnapshotTime *string `field:"required" json:"snapshotTime" yaml:"snapshotTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lightsail_instance#status LightsailInstance#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/lightsail_instance#type LightsailInstance#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

