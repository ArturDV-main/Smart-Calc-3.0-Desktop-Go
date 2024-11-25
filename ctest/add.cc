#include <cstdio>
#include "add.h"
#include <vector>

int add(int a, int b) {
    std::vector<int> v;
    // Add two more integers to vector
    v.push_back(a);
    v.push_back(b);
    return v[0] + v[1];
}
