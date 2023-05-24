import grpc
from . import tracker_pb2 as pb2, tracker_pb2_grpc as pb2_grpc


class TrackerClient:
    def __init__(self):
        self.host = "localhost"
        self.server_port = 3000

        self.channel = grpc.insecure_channel(f"{self.host}:{self.server_port}")
        self.stub = pb2_grpc.TrackedServiceStub(self.channel)

    def create_tracked(
        self, title, description, cover, category, genre, rating, user_id
    ):
        request = pb2.CreateTrackedRequest(
            title=title,
            description=description,
            cover=cover,
            category=category,
            genre=genre,
            rating=rating,
            userID=user_id,
        )
        return self.stub.CreateTracked(request)

    def get_trackeds(self, user_id):
        request = pb2.GetTrackedsRequest(userID=user_id)
        return self.stub.GetTrackeds(request)

    def update_tracked(self, id, title, description, cover, category, genre, rating):
        request = pb2.UpdateTrackedRequest(
            id=id,
            title=title,
            description=description,
            cover=cover,
            category=category,
            genre=genre,
            rating=rating,
        )
        return self.stub.UpdateTracked(request)

    def delete_tracked(self, id):
        request = pb2.DeleteTrackedRequest(id=id)
        return self.stub.DeleteTracked(request)
