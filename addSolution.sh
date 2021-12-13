#!/bin/sh
cp -frp template -T solutions/$1
mv solutions/$1/aoc_day.solution.go solutions/$1/aoc_$1.go
