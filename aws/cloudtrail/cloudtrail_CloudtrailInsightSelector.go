package cloudtrail


type CloudtrailInsightSelector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudtrail#insight_type Cloudtrail#insight_type}.
	InsightType *string `field:"required" json:"insightType" yaml:"insightType"`
}

