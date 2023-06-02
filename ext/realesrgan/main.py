import base64
from io import BytesIO
from basicsr.archs.rrdbnet_arch import RRDBNet
from basicsr.utils.download_util import load_file_from_url
from flask import Flask, request
from realesrgan import RealESRGANer
from realesrgan.archs.srvgg_arch import SRVGGNetCompact
from PIL import Image
import numpy as np
import torch

import logging
import os

app = Flask(__name__)

REAL_ESRGAN_MODELS = {
    "realesr-general-x4v3": {
        "url": "https://github.com/xinntao/Real-ESRGAN/releases/download/v0.2.5.0/realesr-general-x4v3.pth",
        "model_md5": "91a7644643c884ee00737db24e478156",
        "scale": 4,
        "model": lambda: SRVGGNetCompact(
            num_in_ch=3,
            num_out_ch=3,
            num_feat=64,
            num_conv=32,
            upscale=4,
            act_type="prelu",
        ),
    },
    "RealESRGAN_x4plus": {
        "url": "https://github.com/xinntao/Real-ESRGAN/releases/download/v0.1.0/RealESRGAN_x4plus.pth",
        "model_md5": "99ec365d4afad750833258a1a24f44ca",
        "scale": 4,
        "model": lambda: RRDBNet(
            num_in_ch=3,
            num_out_ch=3,
            num_feat=64,
            num_block=23,
            num_grow_ch=32,
            scale=4,
        ),
    },
}

name = "realesr-general-x4v3"

if name not in REAL_ESRGAN_MODELS:
    raise ValueError(f"Unknown RealESRGAN model name: {name}")

model_info = REAL_ESRGAN_MODELS[name]

model_path = os.path.join('weights', name + '.pth')
if not os.path.isfile(model_path):
    ROOT_DIR = os.path.dirname(os.path.abspath(__file__))
    model_path = load_file_from_url(url=model_info["url"], model_dir=os.path.join(ROOT_DIR, 'weights'), progress=True, file_name=None)
logging.info(f"RealESRGAN model path: {model_path}")

device = "cuda" if torch.cuda.is_available() else "cpu"

model = RealESRGANer(
    scale=model_info["scale"],
    model_path=model_path,
    model=model_info["model"](),
    half=True if "cuda" in device else False,
    device=device,
)

@app.route('/superscale', methods=['POST'])
def superscale():
    scale = request.json.get("scale", model_info["scale"])

    image_data = base64.b64decode(request.json['image'])
    image = np.asarray(Image.open(BytesIO(image_data)))
    print(f"RealESRGAN input shape: {image.shape}, scale: {scale}", flush=True)

    upsampled = model.enhance(image, outscale=scale)[0]
    upsampled_img = Image.fromarray(upsampled)
    print(f"RealESRGAN output shape: {upsampled.shape}", flush=True)

    with BytesIO() as buffer:
        upsampled_img.save(buffer, format="jpeg")
        return {"image": base64.b64encode(buffer.getvalue()).decode()}

if __name__ == '__main__':
    print("running")
    app.run(host='0.0.0.0', port=5001)
