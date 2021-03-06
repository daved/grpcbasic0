# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: grpcbasic0.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='grpcbasic0.proto',
  package='pb',
  syntax='proto3',
  serialized_pb=_b('\n\x10grpcbasic0.proto\x12\x02pb\x1a\x1cgoogle/api/annotations.proto\"B\n\x08UserResp\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0b\n\x03\x61ge\x18\x03 \x01(\x03\x12\x0f\n\x07\x66ortune\x18\x04 \x01(\t\"\'\n\nNewUserReq\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0b\n\x03\x61ge\x18\x02 \x01(\x03\"\x18\n\nGetUserReq\x12\n\n\x02id\x18\x01 \x01(\x03\"(\n\tUsersResp\x12\x1b\n\x05users\x18\x01 \x03(\x0b\x32\x0c.pb.UserResp\"9\n\x0bGetUsersReq\x12\r\n\x05start\x18\x01 \x01(\x03\x12\r\n\x05\x63ount\x18\x02 \x01(\x03\x12\x0c\n\x04\x64\x65sc\x18\x03 \x01(\x08\x32\xca\x01\n\x0bUserService\x12<\n\x07NewUser\x12\x0e.pb.NewUserReq\x1a\x0c.pb.UserResp\"\x13\x82\xd3\xe4\x93\x02\r\"\x08/v1/user:\x01*\x12>\n\x07GetUser\x12\x0e.pb.GetUserReq\x1a\x0c.pb.UserResp\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/v1/user/{id}\x12=\n\x08GetUsers\x12\x0f.pb.GetUsersReq\x1a\r.pb.UsersResp\"\x11\x82\xd3\xe4\x93\x02\x0b\x12\t/v1/usersb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_USERRESP = _descriptor.Descriptor(
  name='UserResp',
  full_name='pb.UserResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.UserResp.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='pb.UserResp.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='age', full_name='pb.UserResp.age', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fortune', full_name='pb.UserResp.fortune', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
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
  serialized_start=54,
  serialized_end=120,
)


_NEWUSERREQ = _descriptor.Descriptor(
  name='NewUserReq',
  full_name='pb.NewUserReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='pb.NewUserReq.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='age', full_name='pb.NewUserReq.age', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
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
  serialized_start=122,
  serialized_end=161,
)


_GETUSERREQ = _descriptor.Descriptor(
  name='GetUserReq',
  full_name='pb.GetUserReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='pb.GetUserReq.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
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
  serialized_start=163,
  serialized_end=187,
)


_USERSRESP = _descriptor.Descriptor(
  name='UsersResp',
  full_name='pb.UsersResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='users', full_name='pb.UsersResp.users', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
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
  serialized_start=189,
  serialized_end=229,
)


_GETUSERSREQ = _descriptor.Descriptor(
  name='GetUsersReq',
  full_name='pb.GetUsersReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='start', full_name='pb.GetUsersReq.start', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='count', full_name='pb.GetUsersReq.count', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='desc', full_name='pb.GetUsersReq.desc', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
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
  serialized_start=231,
  serialized_end=288,
)

_USERSRESP.fields_by_name['users'].message_type = _USERRESP
DESCRIPTOR.message_types_by_name['UserResp'] = _USERRESP
DESCRIPTOR.message_types_by_name['NewUserReq'] = _NEWUSERREQ
DESCRIPTOR.message_types_by_name['GetUserReq'] = _GETUSERREQ
DESCRIPTOR.message_types_by_name['UsersResp'] = _USERSRESP
DESCRIPTOR.message_types_by_name['GetUsersReq'] = _GETUSERSREQ
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UserResp = _reflection.GeneratedProtocolMessageType('UserResp', (_message.Message,), dict(
  DESCRIPTOR = _USERRESP,
  __module__ = 'grpcbasic0_pb2'
  # @@protoc_insertion_point(class_scope:pb.UserResp)
  ))
_sym_db.RegisterMessage(UserResp)

NewUserReq = _reflection.GeneratedProtocolMessageType('NewUserReq', (_message.Message,), dict(
  DESCRIPTOR = _NEWUSERREQ,
  __module__ = 'grpcbasic0_pb2'
  # @@protoc_insertion_point(class_scope:pb.NewUserReq)
  ))
_sym_db.RegisterMessage(NewUserReq)

GetUserReq = _reflection.GeneratedProtocolMessageType('GetUserReq', (_message.Message,), dict(
  DESCRIPTOR = _GETUSERREQ,
  __module__ = 'grpcbasic0_pb2'
  # @@protoc_insertion_point(class_scope:pb.GetUserReq)
  ))
_sym_db.RegisterMessage(GetUserReq)

UsersResp = _reflection.GeneratedProtocolMessageType('UsersResp', (_message.Message,), dict(
  DESCRIPTOR = _USERSRESP,
  __module__ = 'grpcbasic0_pb2'
  # @@protoc_insertion_point(class_scope:pb.UsersResp)
  ))
_sym_db.RegisterMessage(UsersResp)

GetUsersReq = _reflection.GeneratedProtocolMessageType('GetUsersReq', (_message.Message,), dict(
  DESCRIPTOR = _GETUSERSREQ,
  __module__ = 'grpcbasic0_pb2'
  # @@protoc_insertion_point(class_scope:pb.GetUsersReq)
  ))
_sym_db.RegisterMessage(GetUsersReq)



_USERSERVICE = _descriptor.ServiceDescriptor(
  name='UserService',
  full_name='pb.UserService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=291,
  serialized_end=493,
  methods=[
  _descriptor.MethodDescriptor(
    name='NewUser',
    full_name='pb.UserService.NewUser',
    index=0,
    containing_service=None,
    input_type=_NEWUSERREQ,
    output_type=_USERRESP,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\r\"\010/v1/user:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='GetUser',
    full_name='pb.UserService.GetUser',
    index=1,
    containing_service=None,
    input_type=_GETUSERREQ,
    output_type=_USERRESP,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\017\022\r/v1/user/{id}')),
  ),
  _descriptor.MethodDescriptor(
    name='GetUsers',
    full_name='pb.UserService.GetUsers',
    index=2,
    containing_service=None,
    input_type=_GETUSERSREQ,
    output_type=_USERSRESP,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\013\022\t/v1/users')),
  ),
])
_sym_db.RegisterServiceDescriptor(_USERSERVICE)

DESCRIPTOR.services_by_name['UserService'] = _USERSERVICE

# @@protoc_insertion_point(module_scope)
