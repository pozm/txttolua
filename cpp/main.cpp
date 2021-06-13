#include <iostream>
#include <fstream>
#include <filesystem>
#define el std::endl
namespace fs = std::filesystem;

int main(int argc, char *argv[] ) {
    fs::path p = argv[1];

    if (!fs::exists(p)) {
        std::cout << "There is no file which exists there";
        return 1;
    }
    if ( p.extension().string() != ".txt" ) {
        std::cout << "sorry you can only convert txt files to lua.";
        return 1;
    }
    fs::path newP = p.parent_path().string() + "\\" + p.stem().string() + ".lua";
    std::cout << "from" << absolute(p) << el << "to" << absolute(newP) << el << el;
    fs::rename(p,newP);
    std::cout << absolute(p) << el;
    return 0;
}
