#!/bin/bash

git clone https://github.com/marmalodak/Aruco_Tracker.git

python3 -m venv --system-site-packages Aruco_Tracker-venv
source Aruco_Tracker-venv/bin/activate
pip3 install --progress-bar pretty --upgrade pip
pip3 install --progress-bar pretty opencv-python opencv-contrib-python icecream

python3 have_numpy.py
have_numpy=$?
if [[ have_numpy -eq 1 ]]; then
    echo Sit back, this might take a while...
    pip install --progress-bar pretty numpy
fi

cd Aruco_Tracker
python3 aruco_tracker.py
