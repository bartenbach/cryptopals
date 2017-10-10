#!/bin/bash

for i in {1..6}; do
  go test -v com/blakebartenbach/cryptopals/challenge$i
done
