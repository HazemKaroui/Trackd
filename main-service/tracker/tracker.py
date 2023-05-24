import json
from flask import Blueprint, g, jsonify, request

from tracker.client import TrackerClient
from middleware.middleware import token_required

tracker_blueprint = Blueprint("tracker_blueprint", __name__)


@tracker_blueprint.route("/create", methods=["POST"])
@token_required
def create_tracked():
    try:
        user_data = g.user_data
        title = request.json.get("title")
        description = request.json.get("description")
        cover = request.json.get("cover")
        category = request.json.get("category")
        genre = request.json.get("genre")
        rating = request.json.get("rating")
        print(user_data, title, description)
        client = TrackerClient()
        result = client.create_tracked(
            title=title,
            description=description,
            cover=cover,
            category=category,
            genre=genre,
            rating=rating,
            user_id=user_data.get("ID"),
        )
        return jsonify({"success": True, "message": result.message})
    except:
        return jsonify({"success": False, "message": result.message})


@tracker_blueprint.route("/get", methods=["GET"])
@token_required
def get_trackeds():
    try:
        user_data = g.user_data
        client = TrackerClient()
        result = client.get_trackeds(user_id=user_data.get("ID"))
        return jsonify(
            {
                "success": True,
                "message": result.message,
                "data": json.loads(result.data),
            }
        )
    except:
        return jsonify({"success": False, "message": result.message, "data": []})


@tracker_blueprint.route("/update", methods=["PUT"])
@token_required
def update_tracked():
    try:
        id = request.json.get("id")
        title = request.json.get("title")
        description = request.json.get("description")
        cover = request.json.get("cover")
        category = request.json.get("category")
        genre = request.json.get("genre")
        rating = request.json.get("rating")
        client = TrackerClient()
        result = client.update_tracked(
            title=title,
            description=description,
            cover=cover,
            category=category,
            genre=genre,
            rating=rating,
            id=id,
        )
        return jsonify({"success": True, "message": result.message})
    except:
        return jsonify({"success": False, "message": result.message, "data": []})


@tracker_blueprint.route("/delete/<int:id>", methods=["DELETE"])
@token_required
def delete_tracked(id):
    try:
        client = TrackerClient()
        result = client.delete_tracked(id)
        return jsonify(
            {
                "success": False,
                "message": result.message,
            }
        )
    except:
        return jsonify(
            {
                "success": False,
                "message": result.message,
            }
        )
