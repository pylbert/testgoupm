#include "testit.hpp"
#include <stdio.h>

using namespace upm;

TESTIT::TESTIT(int bus)
{
    printf("Constructor called with %d\n", bus);
}
int TESTIT::blibbity() {return 666;}
