package apigateway


type ApiGatewayStageCanarySettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#percent_traffic ApiGatewayStage#percent_traffic}.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#stage_variable_overrides ApiGatewayStage#stage_variable_overrides}.
	StageVariableOverrides *map[string]*string `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/api_gateway_stage#use_stage_cache ApiGatewayStage#use_stage_cache}.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

