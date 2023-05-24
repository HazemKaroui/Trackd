from flask import Blueprint, request, jsonify

tracker_blueprint = Blueprint('tracker', __name__)

@tracker_blueprint.route('/', methods=['GET'])
def get_todos():
    return jsonify({
        'message': 'get trackers route'
    })