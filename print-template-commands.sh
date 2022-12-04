#!/bin/bash

day=$1

mkdir day$day
cp -n main_template.txt day$day/main.go
cp -n main_test_template.txt day$day/main_test.go
touch day$day/ex.txt
touch day$day/in.txt
