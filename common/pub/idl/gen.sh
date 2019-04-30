#!/bin/bash

thrift -gen go:package_prefix=github.com/mkzz115/z-service.git/common/pub/idl/gen-go/ $1
