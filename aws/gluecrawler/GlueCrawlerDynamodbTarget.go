package gluecrawler


type GlueCrawlerDynamodbTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/glue_crawler#path GlueCrawler#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/glue_crawler#scan_all GlueCrawler#scan_all}.
	ScanAll interface{} `field:"optional" json:"scanAll" yaml:"scanAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/glue_crawler#scan_rate GlueCrawler#scan_rate}.
	ScanRate *float64 `field:"optional" json:"scanRate" yaml:"scanRate"`
}

