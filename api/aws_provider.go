package main

import (
	"net/http"
	"encoding/json"
)

func ProvidersAWS(w http.ResponseWriter, r *http.Request) {
	resources := []Resource{
		Resource{
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resources)
}
