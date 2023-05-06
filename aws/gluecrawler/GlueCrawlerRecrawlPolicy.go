package gluecrawler


type GlueCrawlerRecrawlPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/glue_crawler#recrawl_behavior GlueCrawler#recrawl_behavior}.
	RecrawlBehavior *string `field:"optional" json:"recrawlBehavior" yaml:"recrawlBehavior"`
}

