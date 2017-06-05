"""
Auto-generated class for VMInfo
"""
from .VMDiskInfo import VMDiskInfo
from .VMNicInfo import VMNicInfo

from . import client_support


class VMInfo(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(cpu, disks, nics):
        """
        :type cpu: list[float]
        :type disks: list[VMDiskInfo]
        :type nics: list[VMNicInfo]
        :rtype: VMInfo
        """

        return VMInfo(
            cpu=cpu,
            disks=disks,
            nics=nics,
        )

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'VMInfo'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'cpu'
        val = data.get(property_name)
        if val is not None:
            datatypes = [float]
            try:
                self.cpu = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'disks'
        val = data.get(property_name)
        if val is not None:
            datatypes = [VMDiskInfo]
            try:
                self.disks = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'nics'
        val = data.get(property_name)
        if val is not None:
            datatypes = [VMNicInfo]
            try:
                self.nics = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
