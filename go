#!/bin/bash

git clone https://github.com/marmalodak/Aruco_Tracker.git

python3 -m venv aruco-play-venv
source aruco-play-venv/bin/activate
pip3 install -r aruco-play/requirements.txt
