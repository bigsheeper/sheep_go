#include "person.h"
#include "cwrap.h"
#include <iostream>

CPerson PersonInit(int age) {
  std::cout << "Person init" << std::endl;
  CPerson cp = new Person(age);
  return (void*)cp;
}

int age(CPerson cp) {
  auto p = (Person*)cp;
  return p->age();
}
