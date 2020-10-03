#! /bin/bash
mkdir .tmp
cd .tmp || exit
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
tar -xf ffmpeg-release-amd64-static.tar.xz --strip-components 1
mv ffmpeg ../
cd ..
rm -rf .tmp
