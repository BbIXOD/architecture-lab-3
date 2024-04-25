#!/bin/bash

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "bgrect 50 750 50 750"
curl -X POST http://localhost:17000 -d "bgrect 10 20 30 40"
curl -X POST http://localhost:17000 -d "figure 500 500"
curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "update"