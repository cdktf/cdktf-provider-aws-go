package sagemakerworkforce


type SagemakerWorkforceSourceIpConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_workforce#cidrs SagemakerWorkforce#cidrs}.
	Cidrs *[]*string `field:"required" json:"cidrs" yaml:"cidrs"`
}

