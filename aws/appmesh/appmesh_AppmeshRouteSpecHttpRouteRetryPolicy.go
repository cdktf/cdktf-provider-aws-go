package appmesh


type AppmeshRouteSpecHttpRouteRetryPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#max_retries AppmeshRoute#max_retries}.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// per_retry_timeout block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#per_retry_timeout AppmeshRoute#per_retry_timeout}
	PerRetryTimeout *AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout `field:"required" json:"perRetryTimeout" yaml:"perRetryTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#http_retry_events AppmeshRoute#http_retry_events}.
	HttpRetryEvents *[]*string `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#tcp_retry_events AppmeshRoute#tcp_retry_events}.
	TcpRetryEvents *[]*string `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
}
