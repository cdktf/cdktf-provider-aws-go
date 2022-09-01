package appstream


type AppstreamStackApplicationSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#enabled AppstreamStack#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#settings_group AppstreamStack#settings_group}.
	SettingsGroup *string `field:"optional" json:"settingsGroup" yaml:"settingsGroup"`
}

