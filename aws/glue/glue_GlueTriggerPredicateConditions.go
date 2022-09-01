package glue


type GlueTriggerPredicateConditions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#crawler_name GlueTrigger#crawler_name}.
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#crawl_state GlueTrigger#crawl_state}.
	CrawlState *string `field:"optional" json:"crawlState" yaml:"crawlState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#job_name GlueTrigger#job_name}.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#logical_operator GlueTrigger#logical_operator}.
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_trigger#state GlueTrigger#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

