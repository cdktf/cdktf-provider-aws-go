package rdsexporttask


type RdsExportTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rds_export_task#create RdsExportTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rds_export_task#delete RdsExportTask#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

