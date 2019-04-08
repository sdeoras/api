package config

import (
	"os"

	"github.com/sdeoras/api/pb"
)

const (
	FunctionName        = "FUNCTION_NAME"
	ProjectId           = "GOOGLE_PROJECT_ID"
	GoogleGCFDomain     = "GOOGLE_GCF_DOMAIN"
	GoogleClientId      = "GOOGLE_CLIENT_ID"
	GoogleClientSecret  = "GOOGLE_CLIENT_SECRET"
	JwtSecretKey        = "JWT_SECRET_KEY"
	CodeLocation        = "CODE_LOCATION"
	CloudFunctionBucket = "CLOUD_FUNCTIONS_BUCKET"
)

// NewConfigFromEnv builds config response from environment variables
func NewConfigFromEnv() *pb.ConfigResponse {
	return &pb.ConfigResponse{
		FuncName:  os.Getenv(FunctionName),
		ProjectId: os.Getenv(ProjectId),
		Domain:    os.Getenv(GoogleGCFDomain),
		Oauth: &pb.OAuthRequest{
			Google: &pb.GoogleOAuthRequest{
				ClientId:     os.Getenv(GoogleClientId),
				ClientSecret: os.Getenv(GoogleClientSecret),
			},
		},
		JwtSecret:    os.Getenv(JwtSecretKey),
		CodeLocation: os.Getenv(CodeLocation),
		BucketName:   os.Getenv(CloudFunctionBucket),
	}
}
