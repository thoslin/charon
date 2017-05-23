# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: group.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.golang.protobuf.ptypes.timestamp import timestamp_pb2 as github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_timestamp_dot_timestamp__pb2
from github.com.golang.protobuf.ptypes.empty import empty_pb2 as github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_empty_dot_empty__pb2
from github.com.golang.protobuf.ptypes.wrappers import wrappers_pb2 as github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_wrappers_dot_wrappers__pb2
from github.com.piotrkowalczuk.qtypes import qtypes_pb2 as github_dot_com_dot_piotrkowalczuk_dot_qtypes_dot_qtypes__pb2
from github.com.piotrkowalczuk.ntypes import ntypes_pb2 as github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='group.proto',
  package='charonrpc',
  syntax='proto3',
  serialized_pb=_b('\n\x0bgroup.proto\x12\tcharonrpc\x1a;github.com/golang/protobuf/ptypes/timestamp/timestamp.proto\x1a\x33github.com/golang/protobuf/ptypes/empty/empty.proto\x1a\x39github.com/golang/protobuf/ptypes/wrappers/wrappers.proto\x1a-github.com/piotrkowalczuk/qtypes/qtypes.proto\x1a-github.com/piotrkowalczuk/ntypes/ntypes.proto\"\xdc\x01\n\x05Group\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12!\n\ncreated_by\x18\x05 \x01(\x0b\x32\r.ntypes.Int64\x12.\n\nupdated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12!\n\nupdated_by\x18\x07 \x01(\x0b\x32\r.ntypes.Int64\"G\n\x12\x43reateGroupRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12#\n\x0b\x64\x65scription\x18\x02 \x01(\x0b\x32\x0e.ntypes.String\"6\n\x13\x43reateGroupResponse\x12\x1f\n\x05group\x18\x01 \x01(\x0b\x32\x10.charonrpc.Group\"\x1d\n\x0fGetGroupRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"3\n\x10GetGroupResponse\x12\x1f\n\x05group\x18\x01 \x01(\x0b\x32\x10.charonrpc.Group\"V\n\x11ListGroupsRequest\x12\x1d\n\x06offset\x18\x64 \x01(\x0b\x32\r.ntypes.Int64\x12\x1c\n\x05limit\x18\x65 \x01(\x0b\x32\r.ntypes.Int64J\x04\x08\x01\x10\x64\"6\n\x12ListGroupsResponse\x12 \n\x06groups\x18\x01 \x03(\x0b\x32\x10.charonrpc.Group\" \n\x12\x44\x65leteGroupRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"c\n\x12ModifyGroupRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x1c\n\x04name\x18\x02 \x01(\x0b\x32\x0e.ntypes.String\x12#\n\x0b\x64\x65scription\x18\x03 \x01(\x0b\x32\x0e.ntypes.String\"6\n\x13ModifyGroupResponse\x12\x1f\n\x05group\x18\x01 \x01(\x0b\x32\x10.charonrpc.Group\"C\n\x1aSetGroupPermissionsRequest\x12\x10\n\x08group_id\x18\x01 \x01(\x03\x12\x13\n\x0bpermissions\x18\x02 \x03(\t\"R\n\x1bSetGroupPermissionsResponse\x12\x0f\n\x07\x63reated\x18\x01 \x01(\x03\x12\x0f\n\x07removed\x18\x02 \x01(\x03\x12\x11\n\tuntouched\x18\x03 \x01(\x03\")\n\x1bListGroupPermissionsRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"3\n\x1cListGroupPermissionsResponse\x12\x13\n\x0bpermissions\x18\x01 \x03(\t2\xbd\x04\n\x0cGroupManager\x12I\n\x06\x43reate\x12\x1d.charonrpc.CreateGroupRequest\x1a\x1e.charonrpc.CreateGroupResponse\"\x00\x12I\n\x06Modify\x12\x1d.charonrpc.ModifyGroupRequest\x1a\x1e.charonrpc.ModifyGroupResponse\"\x00\x12@\n\x03Get\x12\x1a.charonrpc.GetGroupRequest\x1a\x1b.charonrpc.GetGroupResponse\"\x00\x12\x45\n\x04List\x12\x1c.charonrpc.ListGroupsRequest\x1a\x1d.charonrpc.ListGroupsResponse\"\x00\x12\x45\n\x06\x44\x65lete\x12\x1d.charonrpc.DeleteGroupRequest\x1a\x1a.google.protobuf.BoolValue\"\x00\x12\x64\n\x0fListPermissions\x12&.charonrpc.ListGroupPermissionsRequest\x1a\'.charonrpc.ListGroupPermissionsResponse\"\x00\x12\x61\n\x0eSetPermissions\x12%.charonrpc.SetGroupPermissionsRequest\x1a&.charonrpc.SetGroupPermissionsResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_timestamp_dot_timestamp__pb2.DESCRIPTOR,github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_empty_dot_empty__pb2.DESCRIPTOR,github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_wrappers_dot_wrappers__pb2.DESCRIPTOR,github_dot_com_dot_piotrkowalczuk_dot_qtypes_dot_qtypes__pb2.DESCRIPTOR,github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_GROUP = _descriptor.Descriptor(
  name='Group',
  full_name='charonrpc.Group',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.Group.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='charonrpc.Group.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='charonrpc.Group.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='charonrpc.Group.created_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='created_by', full_name='charonrpc.Group.created_by', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='charonrpc.Group.updated_at', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='updated_by', full_name='charonrpc.Group.updated_by', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=294,
  serialized_end=514,
)


_CREATEGROUPREQUEST = _descriptor.Descriptor(
  name='CreateGroupRequest',
  full_name='charonrpc.CreateGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='charonrpc.CreateGroupRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='charonrpc.CreateGroupRequest.description', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=516,
  serialized_end=587,
)


_CREATEGROUPRESPONSE = _descriptor.Descriptor(
  name='CreateGroupResponse',
  full_name='charonrpc.CreateGroupResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group', full_name='charonrpc.CreateGroupResponse.group', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=589,
  serialized_end=643,
)


_GETGROUPREQUEST = _descriptor.Descriptor(
  name='GetGroupRequest',
  full_name='charonrpc.GetGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.GetGroupRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=645,
  serialized_end=674,
)


_GETGROUPRESPONSE = _descriptor.Descriptor(
  name='GetGroupResponse',
  full_name='charonrpc.GetGroupResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group', full_name='charonrpc.GetGroupResponse.group', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=676,
  serialized_end=727,
)


_LISTGROUPSREQUEST = _descriptor.Descriptor(
  name='ListGroupsRequest',
  full_name='charonrpc.ListGroupsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='offset', full_name='charonrpc.ListGroupsRequest.offset', index=0,
      number=100, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='limit', full_name='charonrpc.ListGroupsRequest.limit', index=1,
      number=101, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=729,
  serialized_end=815,
)


_LISTGROUPSRESPONSE = _descriptor.Descriptor(
  name='ListGroupsResponse',
  full_name='charonrpc.ListGroupsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='groups', full_name='charonrpc.ListGroupsResponse.groups', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=817,
  serialized_end=871,
)


_DELETEGROUPREQUEST = _descriptor.Descriptor(
  name='DeleteGroupRequest',
  full_name='charonrpc.DeleteGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.DeleteGroupRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=873,
  serialized_end=905,
)


_MODIFYGROUPREQUEST = _descriptor.Descriptor(
  name='ModifyGroupRequest',
  full_name='charonrpc.ModifyGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.ModifyGroupRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='charonrpc.ModifyGroupRequest.name', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='charonrpc.ModifyGroupRequest.description', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=907,
  serialized_end=1006,
)


_MODIFYGROUPRESPONSE = _descriptor.Descriptor(
  name='ModifyGroupResponse',
  full_name='charonrpc.ModifyGroupResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group', full_name='charonrpc.ModifyGroupResponse.group', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1008,
  serialized_end=1062,
)


_SETGROUPPERMISSIONSREQUEST = _descriptor.Descriptor(
  name='SetGroupPermissionsRequest',
  full_name='charonrpc.SetGroupPermissionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group_id', full_name='charonrpc.SetGroupPermissionsRequest.group_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='permissions', full_name='charonrpc.SetGroupPermissionsRequest.permissions', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1064,
  serialized_end=1131,
)


_SETGROUPPERMISSIONSRESPONSE = _descriptor.Descriptor(
  name='SetGroupPermissionsResponse',
  full_name='charonrpc.SetGroupPermissionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='created', full_name='charonrpc.SetGroupPermissionsResponse.created', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='removed', full_name='charonrpc.SetGroupPermissionsResponse.removed', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='untouched', full_name='charonrpc.SetGroupPermissionsResponse.untouched', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1133,
  serialized_end=1215,
)


_LISTGROUPPERMISSIONSREQUEST = _descriptor.Descriptor(
  name='ListGroupPermissionsRequest',
  full_name='charonrpc.ListGroupPermissionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.ListGroupPermissionsRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1217,
  serialized_end=1258,
)


_LISTGROUPPERMISSIONSRESPONSE = _descriptor.Descriptor(
  name='ListGroupPermissionsResponse',
  full_name='charonrpc.ListGroupPermissionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='permissions', full_name='charonrpc.ListGroupPermissionsResponse.permissions', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1260,
  serialized_end=1311,
)

_GROUP.fields_by_name['created_at'].message_type = github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_timestamp_dot_timestamp__pb2._TIMESTAMP
_GROUP.fields_by_name['created_by'].message_type = github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2._INT64
_GROUP.fields_by_name['updated_at'].message_type = github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_timestamp_dot_timestamp__pb2._TIMESTAMP
_GROUP.fields_by_name['updated_by'].message_type = github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2._INT64
_CREATEGROUPREQUEST.fields_by_name['description'].message_type = github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2._STRING
_CREATEGROUPRESPONSE.fields_by_name['group'].message_type = _GROUP
_GETGROUPRESPONSE.fields_by_name['group'].message_type = _GROUP
_LISTGROUPSREQUEST.fields_by_name['offset'].message_type = github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2._INT64
_LISTGROUPSREQUEST.fields_by_name['limit'].message_type = github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2._INT64
_LISTGROUPSRESPONSE.fields_by_name['groups'].message_type = _GROUP
_MODIFYGROUPREQUEST.fields_by_name['name'].message_type = github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2._STRING
_MODIFYGROUPREQUEST.fields_by_name['description'].message_type = github_dot_com_dot_piotrkowalczuk_dot_ntypes_dot_ntypes__pb2._STRING
_MODIFYGROUPRESPONSE.fields_by_name['group'].message_type = _GROUP
DESCRIPTOR.message_types_by_name['Group'] = _GROUP
DESCRIPTOR.message_types_by_name['CreateGroupRequest'] = _CREATEGROUPREQUEST
DESCRIPTOR.message_types_by_name['CreateGroupResponse'] = _CREATEGROUPRESPONSE
DESCRIPTOR.message_types_by_name['GetGroupRequest'] = _GETGROUPREQUEST
DESCRIPTOR.message_types_by_name['GetGroupResponse'] = _GETGROUPRESPONSE
DESCRIPTOR.message_types_by_name['ListGroupsRequest'] = _LISTGROUPSREQUEST
DESCRIPTOR.message_types_by_name['ListGroupsResponse'] = _LISTGROUPSRESPONSE
DESCRIPTOR.message_types_by_name['DeleteGroupRequest'] = _DELETEGROUPREQUEST
DESCRIPTOR.message_types_by_name['ModifyGroupRequest'] = _MODIFYGROUPREQUEST
DESCRIPTOR.message_types_by_name['ModifyGroupResponse'] = _MODIFYGROUPRESPONSE
DESCRIPTOR.message_types_by_name['SetGroupPermissionsRequest'] = _SETGROUPPERMISSIONSREQUEST
DESCRIPTOR.message_types_by_name['SetGroupPermissionsResponse'] = _SETGROUPPERMISSIONSRESPONSE
DESCRIPTOR.message_types_by_name['ListGroupPermissionsRequest'] = _LISTGROUPPERMISSIONSREQUEST
DESCRIPTOR.message_types_by_name['ListGroupPermissionsResponse'] = _LISTGROUPPERMISSIONSRESPONSE

Group = _reflection.GeneratedProtocolMessageType('Group', (_message.Message,), dict(
  DESCRIPTOR = _GROUP,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.Group)
  ))
_sym_db.RegisterMessage(Group)

CreateGroupRequest = _reflection.GeneratedProtocolMessageType('CreateGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.CreateGroupRequest)
  ))
_sym_db.RegisterMessage(CreateGroupRequest)

CreateGroupResponse = _reflection.GeneratedProtocolMessageType('CreateGroupResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATEGROUPRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.CreateGroupResponse)
  ))
_sym_db.RegisterMessage(CreateGroupResponse)

GetGroupRequest = _reflection.GeneratedProtocolMessageType('GetGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.GetGroupRequest)
  ))
_sym_db.RegisterMessage(GetGroupRequest)

GetGroupResponse = _reflection.GeneratedProtocolMessageType('GetGroupResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETGROUPRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.GetGroupResponse)
  ))
_sym_db.RegisterMessage(GetGroupResponse)

ListGroupsRequest = _reflection.GeneratedProtocolMessageType('ListGroupsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPSREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupsRequest)
  ))
_sym_db.RegisterMessage(ListGroupsRequest)

ListGroupsResponse = _reflection.GeneratedProtocolMessageType('ListGroupsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPSRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupsResponse)
  ))
_sym_db.RegisterMessage(ListGroupsResponse)

DeleteGroupRequest = _reflection.GeneratedProtocolMessageType('DeleteGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.DeleteGroupRequest)
  ))
_sym_db.RegisterMessage(DeleteGroupRequest)

ModifyGroupRequest = _reflection.GeneratedProtocolMessageType('ModifyGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _MODIFYGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ModifyGroupRequest)
  ))
_sym_db.RegisterMessage(ModifyGroupRequest)

ModifyGroupResponse = _reflection.GeneratedProtocolMessageType('ModifyGroupResponse', (_message.Message,), dict(
  DESCRIPTOR = _MODIFYGROUPRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ModifyGroupResponse)
  ))
_sym_db.RegisterMessage(ModifyGroupResponse)

SetGroupPermissionsRequest = _reflection.GeneratedProtocolMessageType('SetGroupPermissionsRequest', (_message.Message,), dict(
  DESCRIPTOR = _SETGROUPPERMISSIONSREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.SetGroupPermissionsRequest)
  ))
_sym_db.RegisterMessage(SetGroupPermissionsRequest)

SetGroupPermissionsResponse = _reflection.GeneratedProtocolMessageType('SetGroupPermissionsResponse', (_message.Message,), dict(
  DESCRIPTOR = _SETGROUPPERMISSIONSRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.SetGroupPermissionsResponse)
  ))
_sym_db.RegisterMessage(SetGroupPermissionsResponse)

ListGroupPermissionsRequest = _reflection.GeneratedProtocolMessageType('ListGroupPermissionsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPPERMISSIONSREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupPermissionsRequest)
  ))
_sym_db.RegisterMessage(ListGroupPermissionsRequest)

ListGroupPermissionsResponse = _reflection.GeneratedProtocolMessageType('ListGroupPermissionsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPPERMISSIONSRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupPermissionsResponse)
  ))
_sym_db.RegisterMessage(ListGroupPermissionsResponse)


try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces


  class GroupManagerStub(object):

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.Create = channel.unary_unary(
          '/charonrpc.GroupManager/Create',
          request_serializer=CreateGroupRequest.SerializeToString,
          response_deserializer=CreateGroupResponse.FromString,
          )
      self.Modify = channel.unary_unary(
          '/charonrpc.GroupManager/Modify',
          request_serializer=ModifyGroupRequest.SerializeToString,
          response_deserializer=ModifyGroupResponse.FromString,
          )
      self.Get = channel.unary_unary(
          '/charonrpc.GroupManager/Get',
          request_serializer=GetGroupRequest.SerializeToString,
          response_deserializer=GetGroupResponse.FromString,
          )
      self.List = channel.unary_unary(
          '/charonrpc.GroupManager/List',
          request_serializer=ListGroupsRequest.SerializeToString,
          response_deserializer=ListGroupsResponse.FromString,
          )
      self.Delete = channel.unary_unary(
          '/charonrpc.GroupManager/Delete',
          request_serializer=DeleteGroupRequest.SerializeToString,
          response_deserializer=github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_wrappers_dot_wrappers__pb2.BoolValue.FromString,
          )
      self.ListPermissions = channel.unary_unary(
          '/charonrpc.GroupManager/ListPermissions',
          request_serializer=ListGroupPermissionsRequest.SerializeToString,
          response_deserializer=ListGroupPermissionsResponse.FromString,
          )
      self.SetPermissions = channel.unary_unary(
          '/charonrpc.GroupManager/SetPermissions',
          request_serializer=SetGroupPermissionsRequest.SerializeToString,
          response_deserializer=SetGroupPermissionsResponse.FromString,
          )


  class GroupManagerServicer(object):

    def Create(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Modify(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def List(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ListPermissions(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def SetPermissions(self, request, context):
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_GroupManagerServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'Create': grpc.unary_unary_rpc_method_handler(
            servicer.Create,
            request_deserializer=CreateGroupRequest.FromString,
            response_serializer=CreateGroupResponse.SerializeToString,
        ),
        'Modify': grpc.unary_unary_rpc_method_handler(
            servicer.Modify,
            request_deserializer=ModifyGroupRequest.FromString,
            response_serializer=ModifyGroupResponse.SerializeToString,
        ),
        'Get': grpc.unary_unary_rpc_method_handler(
            servicer.Get,
            request_deserializer=GetGroupRequest.FromString,
            response_serializer=GetGroupResponse.SerializeToString,
        ),
        'List': grpc.unary_unary_rpc_method_handler(
            servicer.List,
            request_deserializer=ListGroupsRequest.FromString,
            response_serializer=ListGroupsResponse.SerializeToString,
        ),
        'Delete': grpc.unary_unary_rpc_method_handler(
            servicer.Delete,
            request_deserializer=DeleteGroupRequest.FromString,
            response_serializer=github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_wrappers_dot_wrappers__pb2.BoolValue.SerializeToString,
        ),
        'ListPermissions': grpc.unary_unary_rpc_method_handler(
            servicer.ListPermissions,
            request_deserializer=ListGroupPermissionsRequest.FromString,
            response_serializer=ListGroupPermissionsResponse.SerializeToString,
        ),
        'SetPermissions': grpc.unary_unary_rpc_method_handler(
            servicer.SetPermissions,
            request_deserializer=SetGroupPermissionsRequest.FromString,
            response_serializer=SetGroupPermissionsResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'charonrpc.GroupManager', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaGroupManagerServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    def Create(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Modify(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Get(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def List(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Delete(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ListPermissions(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def SetPermissions(self, request, context):
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaGroupManagerStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    def Create(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    Create.future = None
    def Modify(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    Modify.future = None
    def Get(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    Get.future = None
    def List(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    List.future = None
    def Delete(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    Delete.future = None
    def ListPermissions(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    ListPermissions.future = None
    def SetPermissions(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      raise NotImplementedError()
    SetPermissions.future = None


  def beta_create_GroupManager_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('charonrpc.GroupManager', 'Create'): CreateGroupRequest.FromString,
      ('charonrpc.GroupManager', 'Delete'): DeleteGroupRequest.FromString,
      ('charonrpc.GroupManager', 'Get'): GetGroupRequest.FromString,
      ('charonrpc.GroupManager', 'List'): ListGroupsRequest.FromString,
      ('charonrpc.GroupManager', 'ListPermissions'): ListGroupPermissionsRequest.FromString,
      ('charonrpc.GroupManager', 'Modify'): ModifyGroupRequest.FromString,
      ('charonrpc.GroupManager', 'SetPermissions'): SetGroupPermissionsRequest.FromString,
    }
    response_serializers = {
      ('charonrpc.GroupManager', 'Create'): CreateGroupResponse.SerializeToString,
      ('charonrpc.GroupManager', 'Delete'): github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_wrappers_dot_wrappers__pb2.BoolValue.SerializeToString,
      ('charonrpc.GroupManager', 'Get'): GetGroupResponse.SerializeToString,
      ('charonrpc.GroupManager', 'List'): ListGroupsResponse.SerializeToString,
      ('charonrpc.GroupManager', 'ListPermissions'): ListGroupPermissionsResponse.SerializeToString,
      ('charonrpc.GroupManager', 'Modify'): ModifyGroupResponse.SerializeToString,
      ('charonrpc.GroupManager', 'SetPermissions'): SetGroupPermissionsResponse.SerializeToString,
    }
    method_implementations = {
      ('charonrpc.GroupManager', 'Create'): face_utilities.unary_unary_inline(servicer.Create),
      ('charonrpc.GroupManager', 'Delete'): face_utilities.unary_unary_inline(servicer.Delete),
      ('charonrpc.GroupManager', 'Get'): face_utilities.unary_unary_inline(servicer.Get),
      ('charonrpc.GroupManager', 'List'): face_utilities.unary_unary_inline(servicer.List),
      ('charonrpc.GroupManager', 'ListPermissions'): face_utilities.unary_unary_inline(servicer.ListPermissions),
      ('charonrpc.GroupManager', 'Modify'): face_utilities.unary_unary_inline(servicer.Modify),
      ('charonrpc.GroupManager', 'SetPermissions'): face_utilities.unary_unary_inline(servicer.SetPermissions),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_GroupManager_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('charonrpc.GroupManager', 'Create'): CreateGroupRequest.SerializeToString,
      ('charonrpc.GroupManager', 'Delete'): DeleteGroupRequest.SerializeToString,
      ('charonrpc.GroupManager', 'Get'): GetGroupRequest.SerializeToString,
      ('charonrpc.GroupManager', 'List'): ListGroupsRequest.SerializeToString,
      ('charonrpc.GroupManager', 'ListPermissions'): ListGroupPermissionsRequest.SerializeToString,
      ('charonrpc.GroupManager', 'Modify'): ModifyGroupRequest.SerializeToString,
      ('charonrpc.GroupManager', 'SetPermissions'): SetGroupPermissionsRequest.SerializeToString,
    }
    response_deserializers = {
      ('charonrpc.GroupManager', 'Create'): CreateGroupResponse.FromString,
      ('charonrpc.GroupManager', 'Delete'): github_dot_com_dot_golang_dot_protobuf_dot_ptypes_dot_wrappers_dot_wrappers__pb2.BoolValue.FromString,
      ('charonrpc.GroupManager', 'Get'): GetGroupResponse.FromString,
      ('charonrpc.GroupManager', 'List'): ListGroupsResponse.FromString,
      ('charonrpc.GroupManager', 'ListPermissions'): ListGroupPermissionsResponse.FromString,
      ('charonrpc.GroupManager', 'Modify'): ModifyGroupResponse.FromString,
      ('charonrpc.GroupManager', 'SetPermissions'): SetGroupPermissionsResponse.FromString,
    }
    cardinalities = {
      'Create': cardinality.Cardinality.UNARY_UNARY,
      'Delete': cardinality.Cardinality.UNARY_UNARY,
      'Get': cardinality.Cardinality.UNARY_UNARY,
      'List': cardinality.Cardinality.UNARY_UNARY,
      'ListPermissions': cardinality.Cardinality.UNARY_UNARY,
      'Modify': cardinality.Cardinality.UNARY_UNARY,
      'SetPermissions': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'charonrpc.GroupManager', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
