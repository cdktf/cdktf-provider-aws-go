package gluecrawler


type GlueCrawlerSchemaChangePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/glue_crawler#delete_behavior GlueCrawler#delete_behavior}.
	DeleteBehavior *string `field:"optional" json:"deleteBehavior" yaml:"deleteBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/glue_crawler#update_behavior GlueCrawler#update_behavior}.
	UpdateBehavior *string `field:"optional" json:"updateBehavior" yaml:"updateBehavior"`
}

