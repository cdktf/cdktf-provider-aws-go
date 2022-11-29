package dbinstance


type DbInstanceBlueGreenUpdate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_instance#enabled DbInstance#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

