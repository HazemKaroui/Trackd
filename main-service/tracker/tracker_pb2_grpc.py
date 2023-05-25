# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import tracker_pb2 as tracker__pb2


class TrackerServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateTracked = channel.unary_unary(
            "/tracker.TrackerService/CreateTracked",
            request_serializer=tracker__pb2.CreateTrackedRequest.SerializeToString,
            response_deserializer=tracker__pb2.CreateTrackedResponse.FromString,
        )
        self.GetTrackeds = channel.unary_unary(
            "/tracker.TrackerService/GetTrackeds",
            request_serializer=tracker__pb2.GetTrackedsRequest.SerializeToString,
            response_deserializer=tracker__pb2.GetTrackedsResponse.FromString,
        )
        self.UpdateTracked = channel.unary_unary(
            "/tracker.TrackerService/UpdateTracked",
            request_serializer=tracker__pb2.UpdateTrackedRequest.SerializeToString,
            response_deserializer=tracker__pb2.UpdateTrackedResponse.FromString,
        )
        self.DeleteTracked = channel.unary_unary(
            "/tracker.TrackerService/DeleteTracked",
            request_serializer=tracker__pb2.DeleteTrackedRequest.SerializeToString,
            response_deserializer=tracker__pb2.DeleteTrackedResponse.FromString,
        )


class TrackerServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateTracked(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")

    def GetTrackeds(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")

    def UpdateTracked(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")

    def DeleteTracked(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_TrackerServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "CreateTracked": grpc.unary_unary_rpc_method_handler(
            servicer.CreateTracked,
            request_deserializer=tracker__pb2.CreateTrackedRequest.FromString,
            response_serializer=tracker__pb2.CreateTrackedResponse.SerializeToString,
        ),
        "GetTrackeds": grpc.unary_unary_rpc_method_handler(
            servicer.GetTrackeds,
            request_deserializer=tracker__pb2.GetTrackedsRequest.FromString,
            response_serializer=tracker__pb2.GetTrackedsResponse.SerializeToString,
        ),
        "UpdateTracked": grpc.unary_unary_rpc_method_handler(
            servicer.UpdateTracked,
            request_deserializer=tracker__pb2.UpdateTrackedRequest.FromString,
            response_serializer=tracker__pb2.UpdateTrackedResponse.SerializeToString,
        ),
        "DeleteTracked": grpc.unary_unary_rpc_method_handler(
            servicer.DeleteTracked,
            request_deserializer=tracker__pb2.DeleteTrackedRequest.FromString,
            response_serializer=tracker__pb2.DeleteTrackedResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "tracker.TrackerService", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))


# This class is part of an EXPERIMENTAL API.
class TrackerService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateTracked(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/tracker.TrackerService/CreateTracked",
            tracker__pb2.CreateTrackedRequest.SerializeToString,
            tracker__pb2.CreateTrackedResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )

    @staticmethod
    def GetTrackeds(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/tracker.TrackerService/GetTrackeds",
            tracker__pb2.GetTrackedsRequest.SerializeToString,
            tracker__pb2.GetTrackedsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )

    @staticmethod
    def UpdateTracked(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/tracker.TrackerService/UpdateTracked",
            tracker__pb2.UpdateTrackedRequest.SerializeToString,
            tracker__pb2.UpdateTrackedResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )

    @staticmethod
    def DeleteTracked(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/tracker.TrackerService/DeleteTracked",
            tracker__pb2.DeleteTrackedRequest.SerializeToString,
            tracker__pb2.DeleteTrackedResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )