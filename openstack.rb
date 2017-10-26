#!/usr/bin/env ruby

class OpenStackProvider

  def initialize()
    @arguments = {
      'auth_url': {
        'required': true,
        'description': ('The Identity authentication URL. '
                        'If omitted, the OS_AUTH_URL environment '
                        'variable is used.'),
        },
        'region': {
          'required': false,
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
          'required': false,
          'description': ('The Username to login with. If omitted, the '
                          'OS_USERNAME environment variable is used.'),
        },
        'user_id': {
          'required': false,
          'description': ('The User ID to login with. If omitted, the '
                          'OS_USER_ID environment variable is used.'),
        },
        'tenant_id': {
          'required': false,
          'description': ('The ID of the Tenant (Identity v2) or Project'
                          ' (Identity v3) to login with. If omitted, the'
                          ' OS_TENANT_ID or OS_PROJECT_ID environment '
                          'variables are used.'),
        },
        'tenant_name': {
          'required': false,
          'description': ('The Name of the Tenant (Identity v2) or '
                          'Project (Identity v3) to login with. If '
                          'omitted, the OS_TENANT_NAME or '
                          'OS_PROJECT_NAME environment variable are '
                          'used.'),
        },
        'password': {
          'required': false,
          'description': ('The Password to login with. If omitted, the '
                          'OS_PASSWORD environment variable is used.'),
        },
        'token': {
          'required': false,
          'description': ('required if not using user_name and password) A '
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
          'required': false,
          'description': ('The ID of the Domain to scope to (Identity '
                          'v3). If If omitted, the following environment'
                          ' variables are checked (in this order): '
                          'OS_USER_DOMAIN_ID, OS_PROJECT_DOMAIN_ID, '
                          'OS_DOMAIN_ID.)'),
        },
        'domain_name': {
            'required': false,
            'description': ('The Name of the Domain to scope to (Identity '
                            'v3). If omitted, the following environment '
                            'variables are checked (in this order): '
                            'OS_USER_DOMAIN_NAME, OS_PROJECT_DOMAIN_NAME, '
                            'OS_DOMAIN_NAME, DEFAULT_DOMAIN.'),
        },
        'insecure': {
            'required': false,
            'description': ('Trust self-signed SSL certificates. If '
                            'omitted, the OS_INSECURE environment '
                            'variable is used.'),
        },
        'cacert_file': {
          'required': false,
          'description': ('Specify a custom CA certificate when '
                          'communicating over SSL. You can specify '
                          'either a path to the file or the contents'
                          ' of the certificate. If omitted, the '
                          'OS_CACERT environment variable is used.'),
        },
        'cert': {
          'required': false,
          'description': ('Specify client certificate file for SSL '
                          'client authentication. You can specify '
                          'either a path to the file or the contents'
                          ' of the certificate. If omitted the '
                          'OS_CERT environment variable is used.'),
        },
        'key': {
          'required': false,
          'description': ('Specify client private key file for SSL '
                          'client authentication. You can specify '
                          'either a path to the file or the contents'
                          ' of the key. If omitted the OS_KEY '
                          'environment variable is used.'),
        },
        'endpoint_type': {
            'required': false,
            'description': ('Specify which type of endpoint to use from '
                            'the service catalog. It can be set using the '
                            'OS_ENDPOINT_TYPE environment variable. '
                            'If not set, public endpoints is used.'),
        },
        'swauth': {
          'required': false,
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
  end

  def arguments
    @arguments
  end
end