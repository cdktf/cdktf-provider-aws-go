package emrserverlessapplication


type EmrserverlessApplicationAutoStopConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#enabled EmrserverlessApplication#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#idle_timeout_minutes EmrserverlessApplication#idle_timeout_minutes}.
	IdleTimeoutMinutes *float64 `field:"optional" json:"idleTimeoutMinutes" yaml:"idleTimeoutMinutes"`
}

