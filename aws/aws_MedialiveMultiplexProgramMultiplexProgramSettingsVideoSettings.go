// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_multiplex_program#constant_bitrate MedialiveMultiplexProgram#constant_bitrate}.
	ConstantBitrate *float64 `field:"optional" json:"constantBitrate" yaml:"constantBitrate"`
	// statemux_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_multiplex_program#statemux_settings MedialiveMultiplexProgram#statemux_settings}
	StatemuxSettings *MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettings `field:"optional" json:"statemuxSettings" yaml:"statemuxSettings"`
}

