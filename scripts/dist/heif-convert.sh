#!/usr/bin/env bash


echo "Debug: Arguments passed are: Arg1: $1, Arg2: $2, Arg3: $3, Arg4: $4" 1>&2

if [[ -z $3 ]] && [[ -z $4 ]]; then
    echo "Usage: heif-convert <filename> <output>" 1>&2
    exit 1
fi

# Usage: heif-convert [-q quality 0..100] <filename> <output>

if [[ -f "/usr/bin/heif-convert" ]]; then
    /usr/bin/heif-convert -q 92 "$3" "$4"
elif [[  -f "/usr/local/bin/heif-convert" ]]; then
    /usr/local/bin/heif-convert -q 92 "$3" "$4"
else
    echo "heif-convert not found" 1>&2
    exit 1
fi

# Reset Exif orientation flag if output image was rotated based on "QuickTime:Rotation"

if [[ $(/usr/bin/exiftool -n -QuickTime:Rotation "$3") ]]; then
    /usr/bin/exiftool -overwrite_original -P -n '-ModifyDate<FileModifyDate' -Orientation=1 "$4"
else
    /usr/bin/exiftool -overwrite_original -P -n '-ModifyDate<FileModifyDate' "$4"
fi
