#!/bin/bash

find . -name "*.sh" -exec shellcheck {} + 
