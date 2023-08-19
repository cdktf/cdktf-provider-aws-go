package gluecrawler


type GlueCrawlerLineageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/glue_crawler#crawler_lineage_settings GlueCrawler#crawler_lineage_settings}.
	CrawlerLineageSettings *string `field:"optional" json:"crawlerLineageSettings" yaml:"crawlerLineageSettings"`
}

