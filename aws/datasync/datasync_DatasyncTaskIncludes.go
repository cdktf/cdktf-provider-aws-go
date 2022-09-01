package datasync


type DatasyncTaskIncludes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_task#filter_type DatasyncTask#filter_type}.
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_task#value DatasyncTask#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

