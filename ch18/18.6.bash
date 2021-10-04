#!/bin/bash

curl -si "http://localhost:2050/?foo=1&bar=2"
curl -si -X POST -d "some data to send" http://localhost:2050
