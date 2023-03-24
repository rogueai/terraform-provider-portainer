#!/bin/bash

dir="$(pwd)/testacc"

$dir/prism-cli mock --validate-request=false --validate-response=false --check-security=false $dir/swagger.yml &
echo $! > prism.PID

