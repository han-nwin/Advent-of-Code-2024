#include "iostream"
#include "fstream"

//0: LEFT, 1:RIGHT, 2:UP, 3:DOWN
enum {
    LEFT,
    RIGHT,
    UP,
    DOWN,
};

struct Position {
    std::array<int,2> coordinate;
    std::array<std::array<int,2>,4> surrounding;
    char value;

    Position(std::array<int,2>& coor, std::array<std::array<int,2>,4>& surr, char value)
    : coordinate(coor), surrounding(surr), value(value) {};

};

bool is_valid(std::array<int,2> pos, int max_col_idx, int max_row_idx) {
    if (pos[0] < 0 || pos[0] > max_col_idx || pos[1] <0 || pos[1] > max_row_idx) {
        return false;
    } else {
        return true;
    }
}


//Operator overload to print with std::cout
std::ostream& operator<<(std::ostream& os, const Position& position) {
    os << "Cordinate [" << position.coordinate[0] << "," << position.coordinate[1] << "]\n";
    os << "Value: " << position.value << "\n";
    os << "Surrounding [";
    for (int i = 0; i < 4; i++) {
        os << position.surrounding[i][0] << "," << position.surrounding[i][1] << " ";
    }
    os << "]";
    return os;
}

int main(int argc, char* argv[]) {
    //GET INPUT FILE into string []
    if (argc != 2) {
        std::cerr << "Usage: ./main <file-name>" << std::endl;
    }

    std::fstream file(argv[1]);
    if (!file.is_open()) {
        std::cerr << "Failed to open file" << std::endl;
        return -1;
    }

    std::string line;
    std::vector<std::string> lines;
    while (std::getline(file, line)) {
        lines.emplace_back(line);
    }
    file.close();

    for (const auto & line : lines) {
        std::cout << line << std::endl;
    }
    
    std::vector<Position> records;

    for (int l = 0; l < lines.size(); l++) {
        for (int i = 0; i < lines[l].length(); i++) {
            if (lines[l][i] == '0') {
                std::array<int,2> coordinate = {l, i};
                std::array<std::array<int, 2>, 4> surrounding = {{
                    {l, i - 1}, 
                    {l, i + 1}, 
                    {l - 1, i}, 
                    {l + 1, i}
                }};
                records.emplace_back(coordinate,surrounding,'0');
            }
        }
    }

    int max_row_idx = lines.size() - 1;
    int max_col_idx = lines[0].length() - 1;
    std::cout << max_col_idx << " " << max_row_idx << std::endl;

    for (const auto & record : records) {
        std::cout << record << std::endl;
        for (const auto& surr : record.surrounding) {
            if (is_valid(surr, max_col_idx, max_row_idx)) {
                std::cout << "Valid ";
            } else {
                std::cout << "Not-Valid ";
            }
        }
        std::cout << std::endl;
        std::cout << std::endl;
    }

    //NOTE:Implement BFS for path finding


    return 0;
}
