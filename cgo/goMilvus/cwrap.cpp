#include "person.h"
#include "cwrap.h"

CPerson PersonInit(int age) {
  CPerson cp = new Person(age);
  return (void*)cp;
}

int age(CPerson cp) {
  auto p = (Person*)cp;
  return p->age();
}