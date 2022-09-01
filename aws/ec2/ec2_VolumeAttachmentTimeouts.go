package ec2


type VolumeAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/volume_attachment#create VolumeAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/volume_attachment#delete VolumeAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

