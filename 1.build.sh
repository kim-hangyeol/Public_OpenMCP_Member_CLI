#!/bin/bash
apt-get install -y python-pip
pip install -r requirement.txt
pyinstaller -F ketikubecli.py
