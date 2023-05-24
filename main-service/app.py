from flask import Flask, jsonify
from auth.auth import auth_blueprint
from tracker.tracker import tracker_blueprint

app = Flask(__name__)

app.register_blueprint(auth_blueprint, url_prefix="/auth")
app.register_blueprint(tracker_blueprint, url_prefix="/tracker")


@app.route("/", methods=["GET"])
def root_route():
    return jsonify({"message": "Welcome to Trackd"})


if __name__ == "__main__":
    app.run(debug=True)
