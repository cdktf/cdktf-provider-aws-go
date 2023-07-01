package dbinstance


type DbInstanceBlueGreenUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/db_instance#enabled DbInstance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

