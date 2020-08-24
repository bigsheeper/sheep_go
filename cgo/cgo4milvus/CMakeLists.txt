cmake_minimum_required(VERSION 3.16)
project(milvusProject)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC")

set(CMAKE_CXX_STANDARD 14)

add_library(milvusCwrap SHARED
        segment_base.cpp
        cwrap.cpp)

add_executable(main c_main.cpp)
