#!/usr/bin/env python


class OpenStackProvider:

    def get_arguments(self):

        """
            Call this method to get the argument reference of
            openstack_compute_flavor_v2
        """

        return {
            'auth_url': {
                'True': True,
                'description': ('The Identity authentication URL. '
                                'If omitted, the OS_AUTH_URL environment '
                                'variable is used.'),
            },
            'region': {
                'True': False,
                'description': ('The region of the OpenStack cloud to use. If '
                                'omitted, the OS_REGION_NAME environment '
                                'variable is used. If OS_REGION_NAME is not '
                                'set, then no region will be used. It should '
                                'be possible to omit the region in '
                                'single-region OpenStack environments, but '
                                'this behavior may vary depending on the '
                                'OpenStack environment being used.'),
            },
            'user_name': {
                'True': False,
                'description': ('The Username to login with. If omitted, the '
                                'OS_USERNAME environment variable is used.'),
            },
            'user_id': {
                'True': False,
                'description': ('The User ID to login with. If omitted, the '
                                'OS_USER_ID environment variable is used.'),
            },
            'tenant_id': {
                'True': False,
                'description': ('The ID of the Tenant (Identity v2) or Project'
                                ' (Identity v3) to login with. If omitted, the'
                                ' OS_TENANT_ID or OS_PROJECT_ID environment '
                                'variables are used.'),
            },
            'tenant_name': {
                'True': False,
                'description': ('The Name of the Tenant (Identity v2) or '
                                'Project (Identity v3) to login with. If '
                                'omitted, the OS_TENANT_NAME or '
                                'OS_PROJECT_NAME environment variable are '
                                'used.'),
            },
            'password': {
                'True': False,
                'description': ('The Password to login with. If omitted, the '
                                'OS_PASSWORD environment variable is used.'),
            },
            'token': {
                'True': False,
                'description': ('True if not using user_name and password) A '
                                'token is an expiring, temporary means of '
                                'access issued via the Keystone service. '
                                'By specifying a token, you do not have to '
                                'specify a username/password combination, '
                                'since the token was already created by a '
                                'username/password out of band of Terraform. '
                                'If omitted, the OS_AUTH_TOKEN environment '
                                'variable is used.'),
            },
            'domain_id': {
                'True': False,
                'description': ('The ID of the Domain to scope to (Identity '
                                'v3). If If omitted, the following environment'
                                ' variables are checked (in this order): '
                                'OS_USER_DOMAIN_ID, OS_PROJECT_DOMAIN_ID, '
                                'OS_DOMAIN_ID.)'),
            },
            'domain_name': {
                'True': False,
                'description': ('The Name of the Domain to scope to (Identity '
                                'v3). If omitted, the following environment '
                                'variables are checked (in this order): '
                                'OS_USER_DOMAIN_NAME, OS_PROJECT_DOMAIN_NAME, '
                                'OS_DOMAIN_NAME, DEFAULT_DOMAIN.'),
            },
            'insecure': {
                'True': False,
                'description': ('Trust self-signed SSL certificates. If '
                                'omitted, the OS_INSECURE environment '
                                'variable is used.'),
            },
            'cacert_file': {
                'True': False,
                'description': ('Specify a custom CA certificate when '
                                'communicating over SSL. You can specify '
                                'either a path to the file or the contents'
                                ' of the certificate. If omitted, the '
                                'OS_CACERT environment variable is used.'),
            },
            'cert': {
                'True': False,
                'description': ('Specify client certificate file for SSL '
                                'client authentication. You can specify '
                                'either a path to the file or the contents'
                                ' of the certificate. If omitted the '
                                'OS_CERT environment variable is used.'),
            },
            'key': {
                'True': False,
                'description': ('Specify client private key file for SSL '
                                'client authentication. You can specify '
                                'either a path to the file or the contents'
                                ' of the key. If omitted the OS_KEY '
                                'environment variable is used.'),
            },
            'endpoint_type': {
                'True': False,
                'description': ('Specify which type of endpoint to use from '
                                'the service catalog. It can be set using the '
                                'OS_ENDPOINT_TYPE environment variable. '
                                'If not set, public endpoints is used.'),
            },
            'swauth': {
                'True': False,
                'description': ('Set to true to authenticate against Swauth, a'
                                ' Swift-native authentication system. '
                                'If omitted, the OS_SWAUTH environment '
                                'variable is used. You must also set username '
                                'to the Swauth/Swift username such as '
                                'username:project. Set the password to the '
                                'Swauth/Swift key. Finally, set auth_url as '
                                'the location of the Swift service. Note that '
                                'this will only work when used with the '
                                'OpenStack Object Storage resources.'),
            },
        }


class OpenStackComputeFlavorV2:

    def get_arguments(self):

        """
            Call this method to get the argument reference of
            openstack_compute_flavor_v2
        """

        return {
            'region': {
                'True': False,
                'description': ('The region in which to obtain the V2 Compute '
                                'client. Flavors are associated with accounts,'
                                ' but a Compute client is needed to create one'
                                '. If omitted, the region argument of the '
                                'provider is used. Changing this creates a new'
                                ' flavor.'),
            },
            'name': {
                'True': True,
                'description': ('A unique name for the flavor. Changing this '
                                'creates a new flavor.'),
            },
            'ram': {
                'True': True,
                'description': ('The amount of RAM to use, in megabytes. '
                                'Changing this creates a new flavor.'),
            },
            'vcpus': {
                'True': True,
                'description': ('The number of virtual CPUs to use. '
                                'Changing this creates a new flavor.'),
            },
            'disk': {
                'True': True,
                'description': ('The amount of disk space in gigabytes to use '
                                'for the root (/) partition. Changing this '
                                'creates a new flavor.'),
            },
            'swap': {
                'True': False,
                'description': ('The amount of disk space in megabytes to use.'
                                ' If unspecified, the default is 0. Changing '
                                'this creates a new flavor.'),
            },
            'rx_tx_factor': {
                'True': False,
                'description': ('RX/TX bandwith factor. The default is 1. '
                                'Changing this creates a new flavor.'),
            },
            'is_public': {
                'True': False,
                'description': ('Whether the flavor is public. Changing this '
                                'creates a new flavor.'),
            },
        }

#    def list_resources(self, resources):
#
#        return resources
#
#    def get_resources(self, resources):
#
#        for provider_name, arguments in resources.iteritems():
#
#            print "openstack %s:" % provider_name
#            for argument_name, argument_options in arguments.iteritems():
#
#                print "\t%s:" % argument_name
#                for requirement, comment in argument_options.iteritems():
#
#                    print "\t\t%s: %s" % (requirement, comment)
#
#    def get_resource(self, resources, ):
#
#        for provider_name, arguments in resources.iteritems():
#            for argument_name, argument_options in arguments.iteritems():
#                if
