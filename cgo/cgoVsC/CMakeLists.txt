cmake_minimum_required(VERSION 3.16)
project(untitled)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC")

set(CMAKE_CXX_STANDARD 14)

add_library(test SHARED
        test.c)

add_executable(main test.c)
