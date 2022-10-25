package evidentlyproject


type EvidentlyProjectDataDeliveryS3Destination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/evidently_project#bucket EvidentlyProject#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/evidently_project#prefix EvidentlyProject#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

