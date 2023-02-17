package datasynclocationhdfs


type DatasyncLocationHdfsQopConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_location_hdfs#data_transfer_protection DatasyncLocationHdfs#data_transfer_protection}.
	DataTransferProtection *string `field:"optional" json:"dataTransferProtection" yaml:"dataTransferProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_location_hdfs#rpc_protection DatasyncLocationHdfs#rpc_protection}.
	RpcProtection *string `field:"optional" json:"rpcProtection" yaml:"rpcProtection"`
}

