#!/usr/bin/env python

import openstack


class Provider:

    def __init__(self, name, arguments):
        """
            Provider init method
        """
        self.name = name
        self.arguments = arguments


class Resource:

    def __init__(self, type, name, arguments):
        """
            Resource init method
        """
        self.type = type
        self.name = name
        self.arguments = arguments


class Argument:

    def __init__(self, name, required, description):
        """
            Argument init method
        """
        self.name = name
        self.required = required
        self.description = description


def main():

    openstack_provider = Provider(
        name="openstack",
        arguments=openstack.OpenStackProvider().get_arguments()
    )

    openstack_compute_flavor_v2 = Resource(
        type="openstack_compute_flavor_v2",
        name="test",
        arguments=openstack.OpenStackComputeFlavorV2().get_arguments()
    )

    print openstack_provider
    print openstack_compute_flavor_v2
