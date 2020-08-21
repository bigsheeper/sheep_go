#include "person.h"

Person::Person(int age):
  age_(age){}

int Person::age() {
  std::cout << "I'm " << age_ << " years old." << std::endl;
  return age_;
}
