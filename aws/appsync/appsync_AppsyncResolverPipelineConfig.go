package appsync


type AppsyncResolverPipelineConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#functions AppsyncResolver#functions}.
	Functions *[]*string `field:"optional" json:"functions" yaml:"functions"`
}

