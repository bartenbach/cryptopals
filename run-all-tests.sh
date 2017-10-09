#!/bin/bash

for i in {1..3}; do
  go test -v com/blakebartenbach/cryptopals/challenge$i
done
