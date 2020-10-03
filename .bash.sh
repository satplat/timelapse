#! /bin/bash
./ffmpeg -framerate 2 -pattern_type glob -i ".images/*.jpg" -s:v 1440x1080 -c:v libx264 -crf 17 -pix_fmt yuv420p "timelapse.mp4" -y
rm -rf .images