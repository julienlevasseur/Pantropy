package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var awsResources = map[string]Resource{
	"aws_provider": {
		Name: "aws_provider",
		Arguments: []Argument{
			{
				Name:        "access_key",
				Required:    false,
				Description: "This is the AWS access key. It must be provided, but it can also be sourced from the AWS_ACCESS_KEY_ID environment variable, or via a shared credentials file if profile is specified.",
			},
			{
				Name:        "secret_key",
				Required:    false,
				Description: "This is the AWS secret key. It must be provided, but it can also be sourced from the AWS_SECRET_ACCESS_KEY environment variable, or via a shared credentials file if profile is specified.",
			},
			{
				Name:        "region",
				Required:    true,
				Description: "This is the AWS region. It must be provided, but it can also be sourced from the AWS_DEFAULT_REGION environment variables, or via a shared credentials file if profile is specified.",
			},
			{
				Name:        "profile",
				Required:    false,
				Description: "This is the AWS profile name as set in the shared credentials file.",
			},
			{
				Name:        "assume_role",
				Required:    false,
				Description: "An assume_role block (documented below). Only one assume_role block may be in the configuration.",
				Options: []Option{
					{
						Name:        "role_arn",
						Value:       "",
						Required:    true,
						Description: "he ARN of the role to assume.",
					},
					{
						Name:        "session_name",
						Value:       "",
						Required:    false,
						Description: "The session name to use when making the AssumeRole call.",
					},
					{
						Name:        "external_id",
						Value:       "",
						Required:    false,
						Description: "The external ID to use when making the AssumeRole call.",
					},
					{
						Name:        "policy",
						Value:       "",
						Required:    false,
						Description: "A more restrictive policy to apply to the temporary credentials. This gives you a way to further restrict the permissions for the resulting temporary security credentials. You cannot use the passed policy to grant permissions that are in excess of those allowed by the access policy of the role that is being assumed.",
					},
				},
			},
			{
				Name:        "shared_credentials_file",
				Required:    false,
				Description: "This is the path to the shared credentials file. If this is not set and a profile is specified, ~/.aws/credentials will be used.",
			},
			{
				Name:        "token",
				Required:    false,
				Description: "Use this to set an MFA token. It can also be sourced from the AWS_SESSION_TOKEN environment variable.",
			},
			{
				Name:        "max_retries",
				Required:    false,
				Description: "This is the maximum number of times an API call is retried, in the case where requests are being throttled or experiencing transient failures. The delay between the subsequent API calls increases exponentially.",
			},
			{
				Name:        "allowed_account_ids",
				Required:    false,
				Description: "List of allowed, white listed, AWS account IDs to prevent you from mistakenly using an incorrect one (and potentially end up destroying a live environment). Conflicts with forbidden_account_ids.",
			},
			{
				Name:        "forbidden_Account_ids",
				Required:    false,
				Description: "List of forbidden, blacklisted, AWS account IDs to prevent you mistakenly using a wrong one (and potentially end up destroying a live environment). Conflicts with allowed_account_ids.",
			},
			{
				Name:        "insecure",
				Required:    false,
				Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted, default value is false.",
			},
			{
				Name:        "skip_credentials_validation",
				Required:    false,
				Description: "Skip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or implemented.",
			},
			{
				Name:        "skip_get_ec2_platforms",
				Required:    false,
				Description: "Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.",
			},
			{
				Name:        "skip_region_validation",
				Required:    false,
				Description: "Skip validation of provided region name. Useful for AWS-like implementations that use their own region names or to bypass the validation for regions that aren't publicly available yet.",
			},
			{
				Name:        "skip_requesting_account_id",
				Required:    false,
				Description: "Skip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API. When set to true, prevents you from managing any resource that requires Account ID to construct an ARN.",
			},
			{
				Name:        "skip_metadata_api_check",
				Required:    false,
				Description: "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
			},
			{
				Name:        "s3_force_path_style",
				Required:    false,
				Description: "Set this to true to force the request to use path-style addressing, i.e., http://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing, http://BUCKET.s3.amazonaws.com/KEY, when possible. Specific to the Amazon S3 service.",
			},
		},
	},
	"aws_instance": {
		Name: "aws_instance",
		Arguments: []Argument{
			{
				Name:     "ami",
				Required: true,
			},
			{
				Name:     "availability_zone",
				Required: false,
			},
			{
				Name:     "placement_group",
				Required: false,
			},
			{
				Name:     "tenancy",
				Required: false,
			},
			{
				Name:     "ebs_optimized",
				Required: false,
			},
			{
				Name:     "disable_api_termination",
				Required: false,
			},
			{
				Name:     "instance_initiated_shutdown_behavior",
				Required: false,
			},
			{
				Name:     "instance_type",
				Required: true,
			},
			{
				Name:     "key_name",
				Required: false,
			},
			{
				Name:     "monitoring",
				Required: false,
			},
			{
				Name:     "security_groups",
				Required: false,
			},
			{
				Name:     "vpc_security_group_ids",
				Required: false,
			},
			{
				Name:     "subnet_id",
				Required: false,
			},
			{
				Name:     "associate_public_ip_address",
				Required: false,
			},
			{
				Name:     "private_ip",
				Required: false,
			},
			{
				Name:     "source_dest_check",
				Required: false,
			},
			{
				Name:     "user_data",
				Required: false,
			},
			{
				Name:     "iam_instance_profile",
				Required: false,
			},
			{
				Name:     "ipv6_address_count",
				Required: false,
			},
			{
				Name:     "ipv6_addresses",
				Required: false,
			},
			{
				Name:     "tags",
				Required: false,
			},
			{
				Name:     "volume_tags",
				Required: false,
			},
			{
				Name:     "root_block_device",
				Required: false,
				Options: []Option{
					{
						Name:     "volume_type",
						Value:    "standard",
						Required: false,
					},
					{
						Name:     "volume_size",
						Value:    "",
						Required: false,
					},
					{
						Name:     "iops",
						Value:    "",
						Required: false,
					},
					{
						Name:     "delete_on_termination",
						Value:    "true",
						Required: false,
					},
				},
			},
			{
				Name:     "ebs_block_device",
				Required: false,
				Options: []Option{
					{
						Name:     "volume_type",
						Value:    "standard",
						Required: false,
					},
					{
						Name:     "volume_size",
						Value:    "",
						Required: false,
					},
					{
						Name:     "iops",
						Value:    "",
						Required: false,
					},
					{
						Name:     "delete_on_termination",
						Value:    "true",
						Required: false,
					},
					{
						Name:     "device_name",
						Value:    "",
						Required: true,
					},
					{
						Name:     "snapshot_id",
						Value:    "",
						Required: false,
					},
					{
						Name:     "encrypted",
						Value:    "false",
						Required: false,
					},
				},
			},
			{
				Name:     "ephemeral_block_device",
				Required: false,
				Options: []Option{
					{
						Name:     "volume_type",
						Value:    "standard",
						Required: false,
					},
					{
						Name:     "volume_size",
						Value:    "",
						Required: false,
					},
					{
						Name:     "iops",
						Value:    "",
						Required: false,
					},
					{
						Name:     "delete_on_termination",
						Value:    "true",
						Required: false,
					},
					{
						Name:     "device_name",
						Value:    "true",
						Required: true,
					},
					{
						Name:     "virtual_name",
						Value:    "true",
						Required: false,
					},
					{
						Name:     "no_device",
						Value:    "true",
						Required: false,
					},
				},
			},
			{
				Name:     "network_interface",
				Required: false,
				Options: []Option{
					{
						Name:     "device_index",
						Value:    "",
						Required: true,
					},
					{
						Name:     "network_interface_id",
						Value:    "",
						Required: true,
					},
					{
						Name:     "delete_on_termination",
						Value:    "false",
						Required: false,
					},
				},
			},
			{
				Name:     "timeouts",
				Required: false,
				Options: []Option{
					{
						Name:     "create",
						Value:    "10",
						Required: false,
					},
					{
						Name:     "update",
						Value:    "10",
						Required: false,
					},
					{
						Name:     "delete",
						Value:    "10",
						Required: false,
					},
				},
			},
		},
	},
}

// ProvidersAWS : Define the aws provider function
func ProvidersAWS(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(awsResources); err != nil {
		panic(err)
	}
}

// ResourceAWS : Define the aws resource function
func ResourceAWS(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ResourceName := vars["ResourceName"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(awsResources[ResourceName]); err != nil {
		panic(err)
	}
}
