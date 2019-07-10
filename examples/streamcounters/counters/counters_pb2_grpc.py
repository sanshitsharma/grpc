# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import counters_pb2 as counters__pb2


class StatusStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetCounters = channel.unary_stream(
        '/servicestatus.Status/GetCounters',
        request_serializer=counters__pb2.CounterReq.SerializeToString,
        response_deserializer=counters__pb2.CounterResp.FromString,
        )


class StatusServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetCounters(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_StatusServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetCounters': grpc.unary_stream_rpc_method_handler(
          servicer.GetCounters,
          request_deserializer=counters__pb2.CounterReq.FromString,
          response_serializer=counters__pb2.CounterResp.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'servicestatus.Status', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
