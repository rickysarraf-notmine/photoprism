import base64
from flask import Flask, request
from io import BytesIO
from PIL import Image
from ultralytics import YOLO
from ultralytics.yolo.engine.results import Results

app = Flask(__name__)
model = YOLO("yolov8n.pt")

@app.route('/hello', methods=['GET'])
def hello():
    return "elloh"

@app.route('/predict', methods=['POST'])
def predict():
    image_data = base64.b64decode(request.json['image'])
    image = Image.open(BytesIO(image_data))

    results: Results = model.predict(image)

    return results[0].tojson()

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
