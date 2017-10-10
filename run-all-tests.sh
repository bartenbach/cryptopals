#!/bin/bash

for i in {1..4}; do
  go test -v com/blakebartenbach/cryptopals/challenge$i
done
