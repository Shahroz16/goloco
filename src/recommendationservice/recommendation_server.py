import os
import time
from collections import ChainMap
from concurrent import futures

import geopy.distance
import grpc
from grpc_health.v1 import health_pb2_grpc
from grpc_health.v1 import health_pb2

import goloco_pb2
import goloco_pb2_grpc
from logger import get_json_logger

logger = get_json_logger('recommendationservice-server')


def get_distance(starting_location, destination_location):
    coords_1 = (starting_location.latitude, starting_location.longitude)
    coords_2 = (destination_location.latitude, destination_location.longitude)
    return geopy.distance.vincenty(coords_1, coords_2).km


class RecommendationServer(goloco_pb2_grpc.SuggestionServiceServicer):
    def ListSuggestions(self, request, context):
        locations = location_service_stub.GetAllLocations(goloco_pb2.EmptyMessageRequest())
        location_id_map = {}
        for location in locations.location:
            location_id_map[location.id] = location
        # location_id_map = ChainMap(*({location.id: location} for location in locations.location))
        location_ids = (location.id for location in locations.location)
        filtered_locations_ids = list(set(location_ids) - set(request.location_ids))
        recommended_locations = []

        # for requested_id in request.location_ids:
        #     minimum_distance = (
        #         {location_id: get_distance(location_id_map[requested_id], location_id_map[location_id])}
        #         for location_id in filtered_locations_ids)
        #     # for location_id in filtered_locations_ids:
        #     for item in minimum_distance:
        #         logger.info(item)
        #     # minimum_distance = get_distance(location_id_map[requested_id], location_id_map[location_id])
        #     closest_distance_id = ChainMap(*minimum_distance)
        #     logger.info(closest_distance_id)
        #     recommended_locations.append(min(closest_distance_id.keys(), key=(lambda k: closest_distance_id[k])))

        for requested_id in request.location_ids:
            for location_id in filtered_locations_ids:
                logger.info(get_distance(location_id_map.get(requested_id), location_id_map.get(location_id)))
                break
                # minimum_distance = {
                #     location_id: get_distance(location_id_map.get(requested_id), location_id_map.get(location_id))}
                # closest_distance_id = ChainMap(*minimum_distance)
                # recommended_locations.append(min(closest_distance_id.keys(), key=(lambda k: closest_distance_id[k])))

        response = goloco_pb2.ListSuggestionsResponse()
        response.location_ids.extend([location_id_map.get(location_id) for location_id in recommended_locations])

    def Check(self, request, context):
        return health_pb2.HealthCheckResponse(
            status=health_pb2.HealthCheckResponse.SERVING)

    def Watch(self, request, context):
        return health_pb2.HealthCheckResponse(
            status=health_pb2.HealthCheckResponse.SERVING)


if __name__ == '__main__':
    logger.info("initializing recommendation service")
    port = os.environ.get("PORT", "8080")
    location_service_addr = os.environ.get("LOCATION_SERVICE_ADDR")
    if location_service_addr == "":
        raise Exception("LOCATION_SERVICE_ADDR variable not added")
    logger.info("Recommendation Service starting")
    channel = grpc.insecure_channel(location_service_addr)
    location_service_stub = goloco_pb2_grpc.LocationServiceStub(channel)

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service = RecommendationServer()
    goloco_pb2_grpc.add_SuggestionServiceServicer_to_server(service, server)
    health_pb2_grpc.add_HealthServicer_to_server(service, server)

    server.add_insecure_port('[::]:' + port)
    server.start()

    try:
        while True:
            time.sleep(10000)
    except KeyboardInterrupt:
        server.stop()
