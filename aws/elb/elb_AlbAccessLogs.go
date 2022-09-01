package elb


type AlbAccessLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb#bucket Alb#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb#enabled Alb#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb#prefix Alb#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

