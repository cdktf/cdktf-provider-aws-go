package gluecrawler


type GlueCrawlerDeltaTarget struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler#connection_name GlueCrawler#connection_name}.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler#delta_tables GlueCrawler#delta_tables}.
	DeltaTables *[]*string `field:"required" json:"deltaTables" yaml:"deltaTables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_crawler#write_manifest GlueCrawler#write_manifest}.
	WriteManifest interface{} `field:"required" json:"writeManifest" yaml:"writeManifest"`
}

