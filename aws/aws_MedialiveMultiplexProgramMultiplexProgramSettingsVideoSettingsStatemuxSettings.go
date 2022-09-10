// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type MedialiveMultiplexProgramMultiplexProgramSettingsVideoSettingsStatemuxSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_multiplex_program#maximum_bitrate MedialiveMultiplexProgram#maximum_bitrate}.
	MaximumBitrate *float64 `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_multiplex_program#minimum_bitrate MedialiveMultiplexProgram#minimum_bitrate}.
	MinimumBitrate *float64 `field:"optional" json:"minimumBitrate" yaml:"minimumBitrate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/medialive_multiplex_program#priority MedialiveMultiplexProgram#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

