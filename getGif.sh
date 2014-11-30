#!/bin/bash

# clean so not prompted to overwrite
rm -rf output_folder
mkdir output_folder

echo "capturing video"
# get 40 frames of 1920x1080 h264
./capture -F -o -c 40 > ./output_folder/output.raw

echo "converting to gif"
ffmpeg -f h264 -i ./output_folder/output.raw -s 240x135 ./output_folder/output.avi
ffmpeg -i ./output_folder/output.avi -t 8 ./output_folder/snapshot%06d.gif

# gifsicle
gifsicle --delay=10 --loop ./output_folder/*.gif > ./output_folder/animate.gif

# imagemagick
# convert -delay 10 -loop 0 ./output_folder/snapshot*.gif ./output_folder/animate.gif
