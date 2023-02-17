package instance


type InstanceEnclaveOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/instance#enabled Instance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

