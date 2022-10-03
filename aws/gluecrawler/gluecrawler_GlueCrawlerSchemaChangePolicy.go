package gluecrawler


type GlueCrawlerSchemaChangePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler#delete_behavior GlueCrawler#delete_behavior}.
	DeleteBehavior *string `field:"optional" json:"deleteBehavior" yaml:"deleteBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler#update_behavior GlueCrawler#update_behavior}.
	UpdateBehavior *string `field:"optional" json:"updateBehavior" yaml:"updateBehavior"`
}

