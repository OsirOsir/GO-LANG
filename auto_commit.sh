#!/bin/bash

# Add all changes to git
git add .

# Commit changes with a default message
git commit -m "Auto commit $(date)"

# Push changes to the remote repository
git push origin main
