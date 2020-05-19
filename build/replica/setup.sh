#!/bin/bash
sleep 10
echo Starting the replica set
mongo mongodb://mongodb1:27017 replicaSet.js
echo Replica set finished


echo Initializing data
#creating collection here because multi document mongodb transaction can't when error occurs
mongo mongodb://mongodb1:27017,mongodb2:27017,mongodb3:27017/DB?replicaSet=rs0 --eval "db.createCollection(\"applications\");"
echo Data initalized
