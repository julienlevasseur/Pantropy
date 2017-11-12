package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

var openstackResources = map[string]Resource{
	"openstack_compute_flavor_v2": Resource{
		Name: "openstack_compute_flavor_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to obtain the V2 Compute client. Flavors are associated with accounts, but a Compute client is needed to create one. If omitted, the region argument of the provider is used. Changing this creates a new flavor.",
			},
			Argument{
				Name: "name",
				Required: true,
				Description: "A unique name for the flavor. Changing this creates a new flavor.",
			},
			Argument{
				Name: "ram",
				Required: true,
				Description: "The amount of RAM to use, in megabytes. Changing this creates a new flavor.",
			},
			Argument{
				Name: "vcpus",
				Required: true,
				Description: "The number of virtual CPUs to use. Changing this creates a new flavor.",
			},
			Argument{
				Name: "disk",
				Required: true,
				Description: "The amount of disk space in gigabytes to use for the root (/) partition. Changing this creates a new flavor.",
			},
			Argument{
				Name: "swap",
				Required: false,
				Description: "The amount of disk space in megabytes to use. If unspecified, the default is 0. Changing this creates a new flavor.",
			},
			Argument{
				Name: "rx_tx_factor",
				Required: false,
				Description: "RX/TX bandwith factor. The default is 1. Changing this creates a new flavor.",
			},
			Argument{
				Name: "is_public",
				Required: false,
				Description: "Whether the flavor is public. Changing this creates a new flavor.",
			},
		},

	},
	"openstack_compute_floatingip_v2": Resource {
		Name: "openstack_compute_floatingip_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to obtain the V2 Compute client. A Compute client is needed to create a floating IP that can be used with a compute instance. If omitted, the region argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).",
			},
			Argument{
				Name: "pool",
				Required: true,
				Description: "The name of the pool from which to obtain the floating IP. Changing this creates a new floating IP.",
			},
		},

	},
	"openstack_compute_floatingip_associate_v2": Resource {
		Name: "openstack_compute_floatingip_associate_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the region argument of the provider is used. Changing this creates a new floatingip_associate.",
			},
			Argument{
				Name: "floating_ip",
				Required: true,
				Description: "The floating IP to associate.",
			},
			Argument{
				Name: "instance_id",
				Required: true,
				Description: "The instance to associte the floating IP with.",
			},
			Argument{
				Name: "fixed_ip",
				Required: false,
				Description: "The specific IP address to direct traffic to.",
			},
		},
	},
	"openstack_compute_instance_v2": Resource {
		Name: "openstack_compute_instance_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to create the server instance. If omitted, the region argument of the provider is used. Changing this creates a new server.",
			},
			Argument{
				Name: "name",
				Required: true,
				Description: "A unique name for the resource.",
			},
			Argument{
				Name: "image_id",
				Required: false,
				Description: "Required if image_name is empty and not booting from a volume. Do not specify if booting from a volume.) The image ID of the desired image for the server. Changing this creates a new server.",
			},
			Argument{
				Name: "image_name",
				Required: false,
				Description: "Required if image_id is empty and not booting from a volume. Do not specify if booting from a volume.) The name of the desired image for the server. Changing this creates a new server.",
			},
			Argument{
				Name: "flavor_id",
				Required: false,
				Description: "Required if flavor_name is empty) The flavor ID of the desired flavor for the server. Changing this resizes the existing server.",
			},
			Argument{
				Name: "flavor_name",
				Required: false,
				Description: "Required if flavor_id is empty) The name of the desired flavor for the server. Changing this resizes the existing server.",
			},
			Argument{
				Name: "user_data",
				Required: false,
				Description: "The user data to provide when launching the instance. Changing this creates a new server.",
			},
			Argument{
				Name: "security_groups",
				Required: false,
				Description: "An array of one or more security group names to associate with the server. Changing this results in adding/removing security groups from the existing server. Note: When attaching the instance to networks using Ports, place the security groups on the Port and not the instance.",
			},
			Argument{
				Name: "availability_zone",
				Required: false,
				Description: "The availability zone in which to create the server. Changing this creates a new server.",
			},
			Argument{
				Name: "network",
				Required: false,
				Description: "An array of one or more networks to attach to the instance. The network object structure is documented below. Changing this creates a new server.",
				Options: []Option{
					Option{
						Name: "uuid",
						Value: "",
						Required: true,
						Description: "unless port or name is provided) The network UUID to attach to the server. Changing this creates a new server.",
					},
					Option{
						Name: "name",
						Value: "",
						Required: true,
						Description: "unless uuid or port is provided) The human-readable name of the network. Changing this creates a new server.",
					},
					Option{
						Name: "port",
						Value: "",
						Required: true,
						Description: "unless uuid or name is provided) The port UUID of a network to attach to the server. Changing this creates a new server.",
					},
					Option{
						Name: "fixed_ip_v4",
						Value: "",
						Required: false,
						Description: "Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server.",
					},
					Option{
						Name: "fixed_ip_v6",
						Value: "",
						Required: false,
						Description: "Specifies a fixed IPv6 address to be used on this network. Changing this creates a new server.",
					},
					Option{
						Name: "access_network",
						Value: "",
						Required: false,
						Description: "Specifies if this network should be used for provisioning access. Accepts true or false. Defaults to false.",
					},

				},
			},
			Argument{
				Name: "metadata",
				Required: false,
				Description: "Metadata key/value pairs to make available from within the instance. Changing this updates the existing server metadata.",
			},
			Argument{
				Name: "config_drive",
				Required: false,
				Description: "Whether to use the config_drive feature to configure the instance. Changing this creates a new server.",
			},
			Argument{
				Name: "admin_pass",
				Required: false,
				Description: "The administrative password to assign to the server. Changing this changes the root password on the existing server.",
			},
			Argument{
				Name: "key_pair",
				Required: false,
				Description: "The name of a key pair to put on the server. The key pair must already be created and associated with the tenant's account. Changing this creates a new server.",
			},
			Argument{
				Name: "block_device",
				Required: false,
				Description: "Configuration of block devices. The block_device structure is documented below. Changing this creates a new server. You can specify multiple block devices which will create an instance with multiple disks. This configuration is very flexible, so please see the following reference for more information.",
				Options: []Option{
					Option{
						Name: "uuid",
						Value: "",
						Required: true,
						Description: "unless source_type is set to 'blank' ) The UUID of the image, volume, or snapshot. Changing this creates a new server.",
					},
					Option{
						Name: "source_type",
						Value: "",
						Required: true,
						Description: "The source type of the device. Must be one of 'blank', 'image', 'volume', or 'snapshot'. Changing this creates a new server.",
					},
					Option{
						Name: "volume_size",
						Value: "",
						Required: false,
						Description: "The size of the volume to create (in gigabytes). Required in the following combinations: source=image and destination=volume, source=blank and destination=local, and source=blank and destination=volume. Changing this creates a new server.",
					},
					Option{
						Name: "boot_index",
						Value: "",
						Required: false,
						Description: "The boot index of the volume. It defaults to 0. Changing this creates a new server.",
					},
					Option{
						Name: "destination_type",
						Value: "",
						Required: false,
						Description: "The type that gets created. Possible values are 'volume' and 'local'. Changing this creates a new server.",
					},
					Option{
						Name: "delete_on_termination",
						Value: "",
						Required: false,
						Description: "Delete the volume / block device upon termination of the instance. Defaults to false. Changing this creates a new server.",
					},
				},
			},
			Argument{
				Name: "scheduler_hints",
				Required: false,
				Description: "Provide the Nova scheduler with hints on how the instance should be launched. The available hints are described below.",
				Options: []Option{
					Option{
						Name: "group",
						Value: "",
						Required: false,
						Description: "A UUID of a Server Group. The instance will be placed into that group.",
					},
					Option{
						Name: "different_host",
						Value: "",
						Required: false,
						Description: "A list of instance UUIDs. The instance will be scheduled on a different host than all other instances.",
					},
					Option{
						Name: "same_host",
						Value: "",
						Required: false,
						Description: "A list of instance UUIDs. The instance will be scheduled on the same host of those specified.",
					},
					Option{
						Name: "query",
						Value: "",
						Required: false,
						Description: "A conditional query that a compute node must pass in order to host an instance.",
					},
					Option{
						Name: "target_cell",
						Value: "",
						Required: false,
						Description: "The name of a cell to host the instance.",
					},
					Option{
						Name: "build_near_host_ip",
						Value: "",
						Required: false,
						Description: "An IP Address in CIDR form. The instance will be placed on a compute node that is in the same subnet.",
					},
				},
			},
			Argument{
				Name: "personality",
				Required: false,
				Description: "Customize the personality of an instance by defining one or more files and their contents. The personality structure is described below.",
				Options: []Option{
					Option{
						Name: "file",
						Value: "",
						Required: true,
						Description: "The absolute path of the destination file.",
					},
					Option{
						Name: "contents",
						Value: "",
						Required: true,
						Description: "The contents of the file. Limited to 255 bytes.",
					},
				},
			},
			Argument{
				Name: "stop_before_destroy",
				Required: false,
				Description: "Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway.",
			},
			Argument{
				Name: "force_delete",
				Required: false,
				Description: "Whether to force the OpenStack instance to be forcefully deleted. This is useful for environments that have reclaim / soft deletion enabled.",
			},
		},

	},
	"openstack_compute_keypair_v2": Resource {
		Name: "openstack_compute_keypair_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the region argument of the provider is used. Changing this creates a new keypair.",
			},
			Argument{
				Name: "name",
				Required: false,
				Description: "A unique name for the keypair. Changing this creates a new keypair.",
			},
			Argument{
				Name: "public_key",
				Required: true,
				Description: "A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair.",
			},
			Argument{
				Name: "value_specs",
				Required: false,
				Description: "Map of additional options.",
			},
		},
	},
	"openstack_compute_secgroup_v2": Resource {
		Name: "openstack_compute_secgroup_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to obtain the V2 Compute client. A Compute client is needed to create a security group. If omitted, the region argument of the provider is used. Changing this creates a new security group.",
			},
			Argument{
				Name: "name",
				Required: true,
				Description: "A unique name for the security group. Changing this updates the name of an existing security group.",
			},
			Argument{
				Name: "description",
				Required: true,
				Description: "A description for the security group. Changing this updates the description of an existing security group.",
			},
			Argument{
				Name: "rule",
				Required: false,
				Description: "A rule describing how the security group operates. The rule object structure is documented below. Changing this updates the security group rules. As shown in the example above, multiple rule blocks may be used.",
				Options: []Option{
					Option{
						Name: "from_port",
						Value: "",
						Required: true,
						Description: "An integer representing the lower bound of the port range to open. Changing this creates a new security group rule.",
					},
					Option{
						Name: "to_port",
						Value: "",
						Required: true,
						Description: "An integer representing the upper bound of the port range to open. Changing this creates a new security group rule.",
					},
					Option{
						Name: "ip_protocol",
						Value: "",
						Required: true,
						Description: "The protocol type that will be allowed. Changing this creates a new security group rule.",
					},
					Option{
						Name: "cidr",
						Value: "",
						Required: false,
						Description: "Required if from_group_id or self is empty. The IP range that will be the source of network traffic to the security group. Use 0.0.0.0/0 to allow all IP addresses. Changing this creates a new security group rule. Cannot be combined with from_group_id or self.",
					},
					Option{
						Name: "from_group_id",
						Value: "",
						Required: false,
						Description: "Required if cidr or self is empty. The ID of a group from which to forward traffic to the parent group. Changing this creates a new security group rule. Cannot be combined with cidr or self.",
					},
					Option{
						Name: "self",
						Value: "",
						Required: false,
						Description: "Required if cidr and from_group_id is empty. If true, the security group itself will be added as a source to this ingress rule. Cannot be combined with cidr or from_group_id.",
					},

				},
			},
		},

	},
	"openstack_compute_servergroup_v2": Resource {
		Name: "openstack_compute_servergroup_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to obtain the V2 Compute client. If omitted, the region argument of the provider is used. Changing this creates a new server group.",
			},
			Argument{
				Name: "name",
				Required: true,
				Description: "A unique name for the server group. Changing this creates a new server group.",
			},
			Argument{
				Name: "policies",
				Required: true,
				Description: "The set of policies for the server group. Only two two policies are available right now, and both are mutually exclusive. See the Policies section for more information. Changing this creates a new server group.",
			},
			Argument{
				Name: "value_specs",
				Required: false,
				Description: "Map of additional options.",
			},
		},

	},
	"openstack_compute_volume_attach_v2": Resource {
		Name: "openstack_compute_volume_attach_v2",
		Arguments: []Argument{
			Argument{
				Name: "region",
				Required: false,
				Description: "The region in which to obtain the V2 Compute client. A Compute client is needed to create a volume attachment. If omitted, the region argument of the provider is used. Changing this creates a new volume attachment.",
			},
			Argument{
				Name: "instance_id",
				Required: true,
				Description: "The ID of the Instance to attach the Volume to.",
			},
			Argument{
				Name: "volume_id",
				Required: true,
				Description: "The ID of the Volume to attach to an Instance.",
			},
			Argument{
				Name: "device",
				Required: false,
				Description: "The device of the volume attachment (ex: /dev/vdc). NOTE: Being able to specify a device is dependent upon the hypervisor in use. There is a chance that the device specified in Terraform will not be the same device the hypervisor chose. If this happens, Terraform will wish to update the device upon subsequent applying which will cause the volume to be detached and reattached indefinitely. Please use with caution.",
			},
		},
	},
}

// ProvidersOpenstack: Define the openstack provider function
func ProvidersOpenstack(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(openstackResources); err != nil {
		panic(err)
	}
}

// ResourceOpenstack: Define the openstack resource function
func ResourceOpenstack(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ResourceName := vars["ResourceName"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(openstackResources[ResourceName]); err != nil {
		panic(err)
	}
}
