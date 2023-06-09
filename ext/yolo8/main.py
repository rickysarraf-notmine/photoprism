import base64
from flask import Flask, request
from io import BytesIO
from PIL import Image
from typing import List
from ultralytics import YOLO
from ultralytics.yolo.engine.results import Results

app = Flask(__name__)

model = "yolov8n"
detect_model = YOLO(f"{model}.pt")
classify_model = YOLO(f"{model}-cls.pt")

@app.route('/hello', methods=['GET'])
def hello():
    return "elloh"

@app.route('/classify', methods=['POST'])
def classify():
    image_data = base64.b64decode(request.json['image'])
    image = Image.open(BytesIO(image_data))

    results: List[Results] = classify_model.predict(image)
    result = results[0]

    # take only the top3 results
    take = min(len(result.names), 3)
    top_n_idx = result.probs.argsort(0, descending=True)[:take].tolist()

    return {result.names[idx]: result.probs[idx].item() for idx in top_n_idx}

@app.route('/detect', methods=['POST'])
def detect():
    image_data = base64.b64decode(request.json['image'])
    image = Image.open(BytesIO(image_data))

    results: List[Results] = detect_model.predict(image)

    return results[0].tojson()

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
