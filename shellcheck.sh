#!/bin/bash

find . -name "*.sh" -print | xargs -I {} shellcheck {}
