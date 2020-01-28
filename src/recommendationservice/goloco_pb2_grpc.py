# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import goloco_pb2 as goloco__pb2


class LocationServiceStub(object):
  """---------------Location service----------

  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SaveLocation = channel.unary_unary(
        '/goloco.LocationService/SaveLocation',
        request_serializer=goloco__pb2.LocationRequest.SerializeToString,
        response_deserializer=goloco__pb2.LocationResponse.FromString,
        )
    self.GetLocation = channel.unary_unary(
        '/goloco.LocationService/GetLocation',
        request_serializer=goloco__pb2.GetLocationLocationRequest.SerializeToString,
        response_deserializer=goloco__pb2.LocationResponse.FromString,
        )
    self.UpdateLocation = channel.unary_unary(
        '/goloco.LocationService/UpdateLocation',
        request_serializer=goloco__pb2.LocationRequest.SerializeToString,
        response_deserializer=goloco__pb2.LocationResponse.FromString,
        )
    self.DeleteLocation = channel.unary_unary(
        '/goloco.LocationService/DeleteLocation',
        request_serializer=goloco__pb2.DeleteLocationLocationRequest.SerializeToString,
        response_deserializer=goloco__pb2.DeletedLocationId.FromString,
        )
    self.GetAllLocations = channel.unary_unary(
        '/goloco.LocationService/GetAllLocations',
        request_serializer=goloco__pb2.EmptyMessageRequest.SerializeToString,
        response_deserializer=goloco__pb2.AllLocationsResponse.FromString,
        )
    self.GetAllLocationsStream = channel.unary_stream(
        '/goloco.LocationService/GetAllLocationsStream',
        request_serializer=goloco__pb2.EmptyMessageRequest.SerializeToString,
        response_deserializer=goloco__pb2.LocationResponse.FromString,
        )


class LocationServiceServicer(object):
  """---------------Location service----------

  """

  def SaveLocation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetLocation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UpdateLocation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteLocation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetAllLocations(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetAllLocationsStream(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_LocationServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SaveLocation': grpc.unary_unary_rpc_method_handler(
          servicer.SaveLocation,
          request_deserializer=goloco__pb2.LocationRequest.FromString,
          response_serializer=goloco__pb2.LocationResponse.SerializeToString,
      ),
      'GetLocation': grpc.unary_unary_rpc_method_handler(
          servicer.GetLocation,
          request_deserializer=goloco__pb2.GetLocationLocationRequest.FromString,
          response_serializer=goloco__pb2.LocationResponse.SerializeToString,
      ),
      'UpdateLocation': grpc.unary_unary_rpc_method_handler(
          servicer.UpdateLocation,
          request_deserializer=goloco__pb2.LocationRequest.FromString,
          response_serializer=goloco__pb2.LocationResponse.SerializeToString,
      ),
      'DeleteLocation': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteLocation,
          request_deserializer=goloco__pb2.DeleteLocationLocationRequest.FromString,
          response_serializer=goloco__pb2.DeletedLocationId.SerializeToString,
      ),
      'GetAllLocations': grpc.unary_unary_rpc_method_handler(
          servicer.GetAllLocations,
          request_deserializer=goloco__pb2.EmptyMessageRequest.FromString,
          response_serializer=goloco__pb2.AllLocationsResponse.SerializeToString,
      ),
      'GetAllLocationsStream': grpc.unary_stream_rpc_method_handler(
          servicer.GetAllLocationsStream,
          request_deserializer=goloco__pb2.EmptyMessageRequest.FromString,
          response_serializer=goloco__pb2.LocationResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'goloco.LocationService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class SuggestionServiceStub(object):
  """---------------Suggestions service----------

  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListSuggestions = channel.unary_unary(
        '/goloco.SuggestionService/ListSuggestions',
        request_serializer=goloco__pb2.ListSuggestionsRequest.SerializeToString,
        response_deserializer=goloco__pb2.ListSuggestionsResponse.FromString,
        )


class SuggestionServiceServicer(object):
  """---------------Suggestions service----------

  """

  def ListSuggestions(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SuggestionServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListSuggestions': grpc.unary_unary_rpc_method_handler(
          servicer.ListSuggestions,
          request_deserializer=goloco__pb2.ListSuggestionsRequest.FromString,
          response_serializer=goloco__pb2.ListSuggestionsResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'goloco.SuggestionService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class AdServiceStub(object):
  """---------------Ad service----------

  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetAds = channel.unary_unary(
        '/goloco.AdService/GetAds',
        request_serializer=goloco__pb2.AdRequest.SerializeToString,
        response_deserializer=goloco__pb2.AdResponse.FromString,
        )


class AdServiceServicer(object):
  """---------------Ad service----------

  """

  def GetAds(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AdServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetAds': grpc.unary_unary_rpc_method_handler(
          servicer.GetAds,
          request_deserializer=goloco__pb2.AdRequest.FromString,
          response_serializer=goloco__pb2.AdResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'goloco.AdService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))


class SearchServiceStub(object):
  """---------------Search service----------

  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SearchLocation = channel.unary_unary(
        '/goloco.SearchService/SearchLocation',
        request_serializer=goloco__pb2.SearchRequest.SerializeToString,
        response_deserializer=goloco__pb2.SearchResponse.FromString,
        )


class SearchServiceServicer(object):
  """---------------Search service----------

  """

  def SearchLocation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_SearchServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SearchLocation': grpc.unary_unary_rpc_method_handler(
          servicer.SearchLocation,
          request_deserializer=goloco__pb2.SearchRequest.FromString,
          response_serializer=goloco__pb2.SearchResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'goloco.SearchService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))