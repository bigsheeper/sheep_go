cmake_minimum_required(VERSION 3.16)
project(person_project)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fPIC")

set(CMAKE_CXX_STANDARD 14)

add_library(cwrap SHARED
        person.cpp
	cwrap.cpp)

#add_executable(main person.cpp)
