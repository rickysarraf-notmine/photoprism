#!/usr/bin/env bash

TODAY=$(/bin/date -u +%Y%m%d)

MODEL_NAME="NASNet Mobile"
MODEL_URL="https://dl.photoprism.app/tensorflow/nasnet.zip?$TODAY"
MODEL_PATH="assets/nasnet"
MODEL_ZIP="/tmp/photoprism/nasnet.zip"
MODEL_HASH="f18b801354e95cade497b4f12e8d2537d04c04f6  $MODEL_ZIP"
MODEL_VERSION="$MODEL_PATH/version.txt"
MODEL_BACKUP="storage/backup/nasnet-$TODAY"

echo "Installing $MODEL_NAME model for TensorFlow..."

# Create directories
/bin/mkdir -p /tmp/photoprism
/bin/mkdir -p storage/backup

# Check for update
if [[ -f ${MODEL_ZIP} ]] && [[ $(/usr/bin/sha1sum ${MODEL_ZIP}) == "${MODEL_HASH}" ]]; then
  if [[ -f ${MODEL_VERSION} ]]; then
    echo "Already up to date."
    exit
  fi
else
  # Download model
  echo "Downloading latest model from $MODEL_URL..."
  /usr/bin/wget "${MODEL_URL}" -O ${MODEL_ZIP}

  TMP_HASH=$(/usr/bin/sha1sum ${MODEL_ZIP})

  echo "${TMP_HASH}"
fi

# Create backup
if [[ -e ${MODEL_PATH} ]]; then
  echo "Creating backup of existing directory: $MODEL_BACKUP"
  /bin/rm -rf "${MODEL_BACKUP}"
  /bin/mv ${MODEL_PATH} "${MODEL_BACKUP}"
fi

# Unzip model
/usr/bin/unzip ${MODEL_ZIP} -d assets
echo "$MODEL_NAME $TODAY $MODEL_HASH" > ${MODEL_VERSION}

echo "Latest $MODEL_NAME installed."
