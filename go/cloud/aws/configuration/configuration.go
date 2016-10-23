package configuration

//Configuration represents AWS configuration file
type Configuration struct {
	AccessKeyID     string `json:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `json:"AWS_SECRET_ACCESS_KEY"`
	DefaultRegion   string `json:"AWS_DEFAULT_REGION"`
}
