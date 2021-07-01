#!/bin/bash

git clone https://github.com/marmalodak/Aruco_Tracker.git

python3 -m venv Aruco_Tracker-venv
source Aruco_Tracker-venv/bin/activate
pip3 install -r requirements.txt

cd Aruco_Tracker
python aruco_tracker.py
