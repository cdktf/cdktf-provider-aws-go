package datasynctask


type DatasyncTaskExcludes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/datasync_task#filter_type DatasyncTask#filter_type}.
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/datasync_task#value DatasyncTask#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

