# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/event.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.event import event_pb2 as flyteidl_dot_event_dot_event__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/event.proto',
  package='flyteidl.admin',
  syntax='proto3',
  serialized_options=_b('Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin'),
  serialized_pb=_b('\n\x1a\x66lyteidl/admin/event.proto\x12\x0e\x66lyteidl.admin\x1a\x1a\x66lyteidl/event/event.proto\"9\n EventErrorAlreadyInTerminalState\x12\x15\n\rcurrent_phase\x18\x01 \x01(\t\"0\n\x1d\x45ventErrorIncompatibleCluster\x12\x0f\n\x07\x63luster\x18\x01 \x01(\t\"\xc4\x01\n\x12\x45ventFailureReason\x12U\n\x19\x61lready_in_terminal_state\x18\x01 \x01(\x0b\x32\x30.flyteidl.admin.EventErrorAlreadyInTerminalStateH\x00\x12M\n\x14incompatible_cluster\x18\x02 \x01(\x0b\x32-.flyteidl.admin.EventErrorIncompatibleClusterH\x00\x42\x08\n\x06reason\"j\n\x1dWorkflowExecutionEventRequest\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12\x35\n\x05\x65vent\x18\x02 \x01(\x0b\x32&.flyteidl.event.WorkflowExecutionEvent\" \n\x1eWorkflowExecutionEventResponse\"b\n\x19NodeExecutionEventRequest\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12\x31\n\x05\x65vent\x18\x02 \x01(\x0b\x32\".flyteidl.event.NodeExecutionEvent\"\x1c\n\x1aNodeExecutionEventResponse\"b\n\x19TaskExecutionEventRequest\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12\x31\n\x05\x65vent\x18\x02 \x01(\x0b\x32\".flyteidl.event.TaskExecutionEvent\"\x1c\n\x1aTaskExecutionEventResponseB7Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_event_dot_event__pb2.DESCRIPTOR,])




_EVENTERRORALREADYINTERMINALSTATE = _descriptor.Descriptor(
  name='EventErrorAlreadyInTerminalState',
  full_name='flyteidl.admin.EventErrorAlreadyInTerminalState',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='current_phase', full_name='flyteidl.admin.EventErrorAlreadyInTerminalState.current_phase', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=74,
  serialized_end=131,
)


_EVENTERRORINCOMPATIBLECLUSTER = _descriptor.Descriptor(
  name='EventErrorIncompatibleCluster',
  full_name='flyteidl.admin.EventErrorIncompatibleCluster',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cluster', full_name='flyteidl.admin.EventErrorIncompatibleCluster.cluster', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=133,
  serialized_end=181,
)


_EVENTFAILUREREASON = _descriptor.Descriptor(
  name='EventFailureReason',
  full_name='flyteidl.admin.EventFailureReason',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='already_in_terminal_state', full_name='flyteidl.admin.EventFailureReason.already_in_terminal_state', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='incompatible_cluster', full_name='flyteidl.admin.EventFailureReason.incompatible_cluster', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='reason', full_name='flyteidl.admin.EventFailureReason.reason',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=184,
  serialized_end=380,
)


_WORKFLOWEXECUTIONEVENTREQUEST = _descriptor.Descriptor(
  name='WorkflowExecutionEventRequest',
  full_name='flyteidl.admin.WorkflowExecutionEventRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='flyteidl.admin.WorkflowExecutionEventRequest.request_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='event', full_name='flyteidl.admin.WorkflowExecutionEventRequest.event', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=382,
  serialized_end=488,
)


_WORKFLOWEXECUTIONEVENTRESPONSE = _descriptor.Descriptor(
  name='WorkflowExecutionEventResponse',
  full_name='flyteidl.admin.WorkflowExecutionEventResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=490,
  serialized_end=522,
)


_NODEEXECUTIONEVENTREQUEST = _descriptor.Descriptor(
  name='NodeExecutionEventRequest',
  full_name='flyteidl.admin.NodeExecutionEventRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='flyteidl.admin.NodeExecutionEventRequest.request_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='event', full_name='flyteidl.admin.NodeExecutionEventRequest.event', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=524,
  serialized_end=622,
)


_NODEEXECUTIONEVENTRESPONSE = _descriptor.Descriptor(
  name='NodeExecutionEventResponse',
  full_name='flyteidl.admin.NodeExecutionEventResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=624,
  serialized_end=652,
)


_TASKEXECUTIONEVENTREQUEST = _descriptor.Descriptor(
  name='TaskExecutionEventRequest',
  full_name='flyteidl.admin.TaskExecutionEventRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='flyteidl.admin.TaskExecutionEventRequest.request_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='event', full_name='flyteidl.admin.TaskExecutionEventRequest.event', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=654,
  serialized_end=752,
)


_TASKEXECUTIONEVENTRESPONSE = _descriptor.Descriptor(
  name='TaskExecutionEventResponse',
  full_name='flyteidl.admin.TaskExecutionEventResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=754,
  serialized_end=782,
)

_EVENTFAILUREREASON.fields_by_name['already_in_terminal_state'].message_type = _EVENTERRORALREADYINTERMINALSTATE
_EVENTFAILUREREASON.fields_by_name['incompatible_cluster'].message_type = _EVENTERRORINCOMPATIBLECLUSTER
_EVENTFAILUREREASON.oneofs_by_name['reason'].fields.append(
  _EVENTFAILUREREASON.fields_by_name['already_in_terminal_state'])
_EVENTFAILUREREASON.fields_by_name['already_in_terminal_state'].containing_oneof = _EVENTFAILUREREASON.oneofs_by_name['reason']
_EVENTFAILUREREASON.oneofs_by_name['reason'].fields.append(
  _EVENTFAILUREREASON.fields_by_name['incompatible_cluster'])
_EVENTFAILUREREASON.fields_by_name['incompatible_cluster'].containing_oneof = _EVENTFAILUREREASON.oneofs_by_name['reason']
_WORKFLOWEXECUTIONEVENTREQUEST.fields_by_name['event'].message_type = flyteidl_dot_event_dot_event__pb2._WORKFLOWEXECUTIONEVENT
_NODEEXECUTIONEVENTREQUEST.fields_by_name['event'].message_type = flyteidl_dot_event_dot_event__pb2._NODEEXECUTIONEVENT
_TASKEXECUTIONEVENTREQUEST.fields_by_name['event'].message_type = flyteidl_dot_event_dot_event__pb2._TASKEXECUTIONEVENT
DESCRIPTOR.message_types_by_name['EventErrorAlreadyInTerminalState'] = _EVENTERRORALREADYINTERMINALSTATE
DESCRIPTOR.message_types_by_name['EventErrorIncompatibleCluster'] = _EVENTERRORINCOMPATIBLECLUSTER
DESCRIPTOR.message_types_by_name['EventFailureReason'] = _EVENTFAILUREREASON
DESCRIPTOR.message_types_by_name['WorkflowExecutionEventRequest'] = _WORKFLOWEXECUTIONEVENTREQUEST
DESCRIPTOR.message_types_by_name['WorkflowExecutionEventResponse'] = _WORKFLOWEXECUTIONEVENTRESPONSE
DESCRIPTOR.message_types_by_name['NodeExecutionEventRequest'] = _NODEEXECUTIONEVENTREQUEST
DESCRIPTOR.message_types_by_name['NodeExecutionEventResponse'] = _NODEEXECUTIONEVENTRESPONSE
DESCRIPTOR.message_types_by_name['TaskExecutionEventRequest'] = _TASKEXECUTIONEVENTREQUEST
DESCRIPTOR.message_types_by_name['TaskExecutionEventResponse'] = _TASKEXECUTIONEVENTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EventErrorAlreadyInTerminalState = _reflection.GeneratedProtocolMessageType('EventErrorAlreadyInTerminalState', (_message.Message,), dict(
  DESCRIPTOR = _EVENTERRORALREADYINTERMINALSTATE,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.EventErrorAlreadyInTerminalState)
  ))
_sym_db.RegisterMessage(EventErrorAlreadyInTerminalState)

EventErrorIncompatibleCluster = _reflection.GeneratedProtocolMessageType('EventErrorIncompatibleCluster', (_message.Message,), dict(
  DESCRIPTOR = _EVENTERRORINCOMPATIBLECLUSTER,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.EventErrorIncompatibleCluster)
  ))
_sym_db.RegisterMessage(EventErrorIncompatibleCluster)

EventFailureReason = _reflection.GeneratedProtocolMessageType('EventFailureReason', (_message.Message,), dict(
  DESCRIPTOR = _EVENTFAILUREREASON,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.EventFailureReason)
  ))
_sym_db.RegisterMessage(EventFailureReason)

WorkflowExecutionEventRequest = _reflection.GeneratedProtocolMessageType('WorkflowExecutionEventRequest', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWEXECUTIONEVENTREQUEST,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowExecutionEventRequest)
  ))
_sym_db.RegisterMessage(WorkflowExecutionEventRequest)

WorkflowExecutionEventResponse = _reflection.GeneratedProtocolMessageType('WorkflowExecutionEventResponse', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWEXECUTIONEVENTRESPONSE,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowExecutionEventResponse)
  ))
_sym_db.RegisterMessage(WorkflowExecutionEventResponse)

NodeExecutionEventRequest = _reflection.GeneratedProtocolMessageType('NodeExecutionEventRequest', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONEVENTREQUEST,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionEventRequest)
  ))
_sym_db.RegisterMessage(NodeExecutionEventRequest)

NodeExecutionEventResponse = _reflection.GeneratedProtocolMessageType('NodeExecutionEventResponse', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONEVENTRESPONSE,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionEventResponse)
  ))
_sym_db.RegisterMessage(NodeExecutionEventResponse)

TaskExecutionEventRequest = _reflection.GeneratedProtocolMessageType('TaskExecutionEventRequest', (_message.Message,), dict(
  DESCRIPTOR = _TASKEXECUTIONEVENTREQUEST,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskExecutionEventRequest)
  ))
_sym_db.RegisterMessage(TaskExecutionEventRequest)

TaskExecutionEventResponse = _reflection.GeneratedProtocolMessageType('TaskExecutionEventResponse', (_message.Message,), dict(
  DESCRIPTOR = _TASKEXECUTIONEVENTRESPONSE,
  __module__ = 'flyteidl.admin.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskExecutionEventResponse)
  ))
_sym_db.RegisterMessage(TaskExecutionEventResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
