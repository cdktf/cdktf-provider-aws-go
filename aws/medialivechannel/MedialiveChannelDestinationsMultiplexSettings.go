package medialivechannel


type MedialiveChannelDestinationsMultiplexSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#multiplex_id MedialiveChannel#multiplex_id}.
	MultiplexId *string `field:"required" json:"multiplexId" yaml:"multiplexId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_channel#program_name MedialiveChannel#program_name}.
	ProgramName *string `field:"required" json:"programName" yaml:"programName"`
}

