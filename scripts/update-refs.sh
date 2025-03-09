#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

cp "$HOME/.vine/.auth/vine_run_auth.json" "$PROJECT_ROOT/config/"

# Copy .bashrc
startDelimiter="# Vine stuff"
endDelimiter="# ffuts eniV"

bashrcTemplateFile="$PROJECT_ROOT/config/template_files/.bashrc"

> "$bashrcTemplateFile"
startDelimiterFound=false

while IFS= read -r line; do
    if [[ $line = $startDelimiter ]]; then
        startDelimiterFound=true    
    fi
    if $startDelimiterFound; then
        echo "$line" >> "$bashrcTemplateFile"
    fi
    if [[ $line = $endDelimiter ]]; then
        break
    fi
done < $HOME/.bashrc

