#!/bin/bash
#Display the name, power, and gender of the hero using a here document

curl -s https://platform.zone01.gr/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender' 