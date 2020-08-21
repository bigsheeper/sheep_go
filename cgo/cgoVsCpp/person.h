#pragma once

#include <iostream>

class Person {
public:
    explicit Person(int age);

    int age();

private:
    int age_;
};
