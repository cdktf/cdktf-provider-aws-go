package appmesh


type AppmeshRouteSpecHttpRoute struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#action AppmeshRoute#action}
	Action *AppmeshRouteSpecHttpRouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttpRouteMatch `field:"required" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#retry_policy AppmeshRoute#retry_policy}
	RetryPolicy *AppmeshRouteSpecHttpRouteRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appmesh_route#timeout AppmeshRoute#timeout}
	Timeout *AppmeshRouteSpecHttpRouteTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

