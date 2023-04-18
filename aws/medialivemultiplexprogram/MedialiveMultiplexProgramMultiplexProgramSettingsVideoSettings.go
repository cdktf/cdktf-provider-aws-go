package medialivemultiplexprogram


type MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_multiplex_program#constant_bitrate MedialiveMultiplexProgram#constant_bitrate}.
	ConstantBitrate *float64 `field:"optional" json:"constantBitrate" yaml:"constantBitrate"`
	// statemux_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_multiplex_program#statemux_settings MedialiveMultiplexProgram#statemux_settings}
	StatemuxSettings interface{} `field:"optional" json:"statemuxSettings" yaml:"statemuxSettings"`
	// statmux_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/medialive_multiplex_program#statmux_settings MedialiveMultiplexProgram#statmux_settings}
	StatmuxSettings interface{} `field:"optional" json:"statmuxSettings" yaml:"statmuxSettings"`
}

