package lambdafunction


type LambdaEventSourceMappingSelfManagedEventSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#endpoints LambdaEventSourceMapping#endpoints}.
	Endpoints *map[string]*string `field:"required" json:"endpoints" yaml:"endpoints"`
}

