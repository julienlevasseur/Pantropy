package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var awsResources = map[string]Resource{
	"aws_provider": Resource{
		Name: "aws_provider",
		Arguments: []Argument{
			Argument{
				Name: "access_key",
				Required: false,
				Description: "This is the AWS access key. It must be provided, but it can also be sourced from the AWS_ACCESS_KEY_ID environment variable, or via a shared credentials file if profile is specified.",
			},
			Argument{
				Name: "secret_key",
				Required: false,
				Description: "This is the AWS secret key. It must be provided, but it can also be sourced from the AWS_SECRET_ACCESS_KEY environment variable, or via a shared credentials file if profile is specified.",
			},
			Argument{
				Name: "region",
				Required: true,
				Description: "This is the AWS region. It must be provided, but it can also be sourced from the AWS_DEFAULT_REGION environment variables, or via a shared credentials file if profile is specified.",
			},
			Argument{
				Name: "profile",
				Required: false,
				Description: "This is the AWS profile name as set in the shared credentials file.",
			},
			Argument{
				Name: "assume_role",
				Required: false,
				Description: "An assume_role block (documented below). Only one assume_role block may be in the configuration.",
				Options: []Option{
					Option{
						Name: "role_arn",
						Value: "",
						Required: true,
						Description: "he ARN of the role to assume.",
					},
					Option{
						Name: "session_name",
						Value: "",
						Required: false,
						Description: "The session name to use when making the AssumeRole call.",
					},
					Option{
						Name: "external_id",
						Value: "",
						Required: false,
						Description: "The external ID to use when making the AssumeRole call.",
					},
					Option{
						Name: "policy",
						Value: "",
						Required: false,
						Description: "A more restrictive policy to apply to the temporary credentials. This gives you a way to further restrict the permissions for the resulting temporary security credentials. You cannot use the passed policy to grant permissions that are in excess of those allowed by the access policy of the role that is being assumed.",
					},
				},
			},
			Argument{
				Name: "shared_credentials_file",
				Required: false,
				Description: "This is the path to the shared credentials file. If this is not set and a profile is specified, ~/.aws/credentials will be used.",
			},
			Argument{
				Name: "token",
				Required: false,
				Description: "Use this to set an MFA token. It can also be sourced from the AWS_SESSION_TOKEN environment variable.",
			},
			Argument{
				Name: "max_retries",
				Required: false,
				Description: "This is the maximum number of times an API call is retried, in the case where requests are being throttled or experiencing transient failures. The delay between the subsequent API calls increases exponentially.",
			},
			Argument{
				Name: "allowed_account_ids",
				Required: false,
				Description: "List of allowed, white listed, AWS account IDs to prevent you from mistakenly using an incorrect one (and potentially end up destroying a live environment). Conflicts with forbidden_account_ids.",
			},
			Argument{
				Name: "forbidden_Account_ids",
				Required: false,
				Description: "List of forbidden, blacklisted, AWS account IDs to prevent you mistakenly using a wrong one (and potentially end up destroying a live environment). Conflicts with allowed_account_ids.",
			},
			Argument{
				Name: "insecure",
				Required: false,
				Description: "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted, default value is false.",
			},
			Argument{
				Name: "skip_credentials_validation",
				Required: false,
				Description: "Skip the credentials validation via the STS API. Useful for AWS API implementations that do not have STS available or implemented.",
			},
			Argument{
				Name: "skip_get_ec2_platforms",
				Required: false,
				Description: "Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.",
			},
			Argument{
				Name: "skip_region_validation",
				Required: false,
				Description: "Skip validation of provided region name. Useful for AWS-like implementations that use their own region names or to bypass the validation for regions that aren't publicly available yet.",
			},
			Argument{
				Name: "skip_requesting_account_id",
				Required: false,
				Description: "Skip requesting the account ID. Useful for AWS API implementations that do not have the IAM, STS API, or metadata API. When set to true, prevents you from managing any resource that requires Account ID to construct an ARN.",
			},
			Argument{
				Name: "skip_metadata_api_check",
				Required: false,
				Description: "Skip the AWS Metadata API check. Useful for AWS API implementations that do not have a metadata API endpoint. Setting to true prevents Terraform from authenticating via the Metadata API. You may need to use other authentication methods like static credentials, configuration variables, or environment variables.",
			},
			Argument{
				Name: "s3_force_path_style",
				Required: false,
				Description: "Set this to true to force the request to use path-style addressing, i.e., http://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing, http://BUCKET.s3.amazonaws.com/KEY, when possible. Specific to the Amazon S3 service.",
			},
		},
	},
	"aws_instance": Resource{
		Name: "aws_instance",
		Arguments: []Argument{
			Argument{
				Name: "ami",
				Required: true,
			},
			Argument{
				Name: "availability_zone",
				Required: false,
			},
			Argument{
				Name: "placement_group",
				Required: false,
			},
			Argument{
				Name: "tenancy",
				Required: false,
			},
			Argument{
				Name: "ebs_optimized",
				Required: false,
			},
			Argument{
				Name: "disable_api_termination",
				Required: false,
			},
			Argument{
				Name: "instance_initiated_shutdown_behavior",
				Required: false,
			},
			Argument{
				Name: "instance_type",
				Required: true,
			},
			Argument{
				Name: "key_name",
				Required: false,
			},
			Argument{
				Name: "monitoring",
				Required: false,
			},
			Argument{
				Name: "security_groups",
				Required: false,
			},
			Argument{
				Name: "vpc_security_group_ids",
				Required: false,
			},
			Argument{
				Name: "subnet_id",
				Required: false,
			},
			Argument{
				Name: "associate_public_ip_address",
				Required: false,
			},
			Argument{
				Name: "private_ip",
				Required: false,
			},
			Argument{
				Name: "source_dest_check",
				Required: false,
			},
			Argument{
				Name: "user_data",
				Required: false,
			},
			Argument{
				Name: "iam_instance_profile",
				Required: false,
			},
			Argument{
				Name: "ipv6_address_count",
				Required: false,
			},
			Argument{
				Name: "ipv6_addresses",
				Required: false,
			},
			Argument{
				Name: "tags",
				Required: false,
			},
			Argument{
				Name: "volume_tags",
				Required: false,
			},
			Argument{
				Name: "root_block_device",
				Required: false,
				Options: []Option{
					Option{
						Name: "volume_type",
						Value: "standard",
						Required: false,
					},
					Option{
						Name: "volume_size",
						Value: "",
						Required: false,
					},
					Option{
						Name: "iops",
						Value: "",
						Required: false,
					},
					Option{
						Name: "delete_on_termination",
						Value: "true",
						Required: false,
					},
				},
			},
			Argument{
				Name: "ebs_block_device",
				Required: false,
				Options: []Option{
					Option{
						Name: "volume_type",
						Value: "standard",
						Required: false,
					},
					Option{
						Name: "volume_size",
						Value: "",
						Required: false,
					},
					Option{
						Name: "iops",
						Value: "",
						Required: false,
					},
					Option{
						Name: "delete_on_termination",
						Value: "true",
						Required: false,
					},
					Option{
						Name: "device_name",
						Value: "",
						Required: true,
					},
					Option{
						Name: "snapshot_id",
						Value: "",
						Required: false,
					},
					Option{
						Name: "encrypted",
						Value: "false",
						Required: false,
					},
				},
			},
			Argument{
				Name: "ephemeral_block_device",
				Required: false,
				Options: []Option{
					Option{
						Name: "volume_type",
						Value: "standard",
						Required: false,
					},
					Option{
						Name: "volume_size",
						Value: "",
						Required: false,
					},
					Option{
						Name: "iops",
						Value: "",
						Required: false,
					},
					Option{
						Name: "delete_on_termination",
						Value: "true",
						Required: false,
					},
					Option{
						Name: "device_name",
						Value: "true",
						Required: true,
					},
					Option{
						Name: "virtual_name",
						Value: "true",
						Required: false,
					},
					Option{
						Name: "no_device",
						Value: "true",
						Required: false,
					},
				},
			},
			Argument{
				Name: "network_interface",
				Required: false,
				Options: []Option{
					Option{
						Name: "device_index",
						Value: "",
						Required: true,
					},
					Option{
						Name: "network_interface_id",
						Value: "",
						Required: true,
					},
					Option{
						Name: "delete_on_termination",
						Value: "false",
						Required: false,
					},
				},
			},
			Argument{
				Name: "timeouts",
				Required: false,
				Options: []Option{
					Option{
						Name: "create",
						Value: "10",
						Required: false,
					},
					Option{
						Name: "update",
						Value: "10",
						Required: false,
					},
					Option{
						Name: "delete",
						Value: "10",
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
