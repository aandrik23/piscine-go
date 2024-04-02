#!/bin/bash
#Display the name of who you are using a here document

curl -s https://platform.zone01.gr/assets/superhero/all.json | jq' .[] | select(.id == 70) | .name'


