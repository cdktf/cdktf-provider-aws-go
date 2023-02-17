package medialivechannel


type MedialiveChannelCdiInputSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#resolution MedialiveChannel#resolution}.
	Resolution *string `field:"required" json:"resolution" yaml:"resolution"`
}

