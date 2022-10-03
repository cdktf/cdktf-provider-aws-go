package inspectorassessmenttemplate


type InspectorAssessmentTemplateEventSubscription struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/inspector_assessment_template#event InspectorAssessmentTemplate#event}.
	Event *string `field:"required" json:"event" yaml:"event"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/inspector_assessment_template#topic_arn InspectorAssessmentTemplate#topic_arn}.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

