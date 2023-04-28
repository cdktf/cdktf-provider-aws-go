package instance


type InstanceEnclaveOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/instance#enabled Instance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

