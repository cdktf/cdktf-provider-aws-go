package rdsexporttask


type RdsExportTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/rds_export_task#create RdsExportTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/rds_export_task#delete RdsExportTask#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

