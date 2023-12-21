#include <iostream>
#include <random>
#include <string>

using namespace std;

string genkey() {
    string lcnkey = "";
    for (int anu = 0; anu < 4; anu++) {
        string block = "";
        for (int j = 0; j < 4; j++) {
            string charSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
            char randomChar = charSet[rand() % charSet.length()];
            block += randomChar;
        }
        lcnkey += block + "-";
    }
    lcnkey = lcnkey.substr(0, lcnkey.length() - 1);
    return lcnkey;
}

int main() {
    cout << genkey() << endl;
    return 0;
}
