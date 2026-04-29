#!/usr/bin/env fish

for file in (rg -l '^apiVersion')
    kubectl apply -f $file
end
