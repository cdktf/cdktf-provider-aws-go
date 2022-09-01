package glue


type GlueCrawlerLineageConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler#crawler_lineage_settings GlueCrawler#crawler_lineage_settings}.
	CrawlerLineageSettings *string `field:"optional" json:"crawlerLineageSettings" yaml:"crawlerLineageSettings"`
}

