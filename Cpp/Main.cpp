#include <string>
#include <algorithm>
#include <random>
 
using namespace std;
 
string random_text(int length) {
    string text;
 
    generate_n(text.begin(), length, []() -> char {
        random_device rnd;
        mt19937 mt(rnd());
        const char char_set[] = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
        char result = char_set[mt() % (sizeof(char_set) - 1)];
        return result;
    });
 
    return text;
}
 
int main(void){
    const int length = 10;
    printf("%s", random_text(length).c_str());
}