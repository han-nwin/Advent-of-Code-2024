#include <iostream>
#include <fstream>
#include <unordered_set>
#include <array>
#include <vector>
#include <unordered_map>
#include <queue>

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
    
    //Default constructor
    Position() : coordinate{{-1,-1}}, 
        surrounding{{{-1,-1}, {-1,-1}, {-1, -1}, {-1, -1}}},
        value(' ') {};

    Position(std::array<int,2>& coor, std::array<std::array<int,2>,4>& surr, char value)
    : coordinate(coor), surrounding(surr), value(value) {};

};

//Custom hash funtion for std::array<int,2>
struct ArrayHasher {
std::size_t operator()(const std::array<int,2>& arr) const {
        //Combine the two elements in the array to create a hash
        //using XOR ^ operator
        return std::hash<int>()(arr[0]) ^ (std::hash<int>()(arr[1]) << 1);
    }
};

//Custom == operator for std::array<int,2>
struct ArrayEqual {
    bool operator()(const std::array<int,2>& lhs, const std::array<int,2>& rhs) const {
        return lhs[0] == rhs[0] && lhs[1] == rhs[1];
    }
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
    
    std::unordered_map<std::array<int,2>,Position, ArrayHasher, ArrayEqual> records;//Store all positions
    std::vector<Position> starts;//Store starting positions

    for (int l = 0; l < lines.size(); l++) {
        for (int i = 0; i < lines[l].length(); i++) {
            std::array<int,2> coordinate = {l, i};
            std::array<std::array<int, 2>, 4> surrounding = {{
                {l, i - 1}, 
                {l, i + 1}, 
                {l - 1, i}, 
                {l + 1, i}
            }};
            char value = lines[l][i];
            Position new_record{coordinate, surrounding, value};
            records[coordinate] = new_record;
            if (value == '0') {
                starts.emplace_back(coordinate,surrounding,'0');
            }
            
        }
    }

    int max_row_idx = lines.size() - 1;
    int max_col_idx = lines[0].length() - 1;

    //NOTE:Implement BFS for path finding

    int ans1 = 0;

    for (const Position& start : starts) {
        std::queue<Position> queue;
        queue.push(start);

        bool flag = false;
        int count = 0;

        std::unordered_set<std::array<int,2>, ArrayHasher, ArrayEqual> visited;

        while (!queue.empty()) {
            Position pos = queue.front();
            queue.pop();

            //Skip visited position
            if (visited.find(pos.coordinate) != visited.end()) {
                continue;
            }
            visited.insert(pos.coordinate);

            if (pos.value == '9') {
                flag = true;
                count++;
                continue;
            }
            
            //Validate the surrounding
            if (is_valid(pos.surrounding[LEFT], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[LEFT]].value == (pos.value + 1)){
                queue.push(records[pos.surrounding[LEFT]]); //Push the element in the queue
            } 
            if (is_valid(pos.surrounding[RIGHT], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[RIGHT]].value == (pos.value + 1)){
                queue.push(records[pos.surrounding[RIGHT]]); //Push the element in the queue
            } 
            if (is_valid(pos.surrounding[UP], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[UP]].value == (pos.value + 1)){
                queue.push(records[pos.surrounding[UP]]); //Push the element in the queue
            } 
            if (is_valid(pos.surrounding[DOWN], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[DOWN]].value == (pos.value + 1)){
                queue.push(records[pos.surrounding[DOWN]]); //Push the element in the queue
            } 
        }

        if (flag) {
            ans1 += count;
        }

    }

    std::cout << "Answer 1: " << ans1 << std::endl;

    //NOTE: PART 2:

    int ans2 = 0;

    for (const Position& start : starts) {
        //Pair of Position and it's coordinate. Use coordinate to make unique path later
        std::queue<std::pair<Position, std::unordered_set<std::array<int,2>, ArrayHasher, ArrayEqual>>> queue;

        queue.push({start, {start.coordinate}});

        bool flag = false;
        int count = 0;

        while (!queue.empty()) {
            auto[pos, path] = queue.front();
            queue.pop();


            if (pos.value == '9') {
                flag = true;
                count++;
                continue;
            }
            
            //Validate the surrounding
            if (is_valid(pos.surrounding[LEFT], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[LEFT]].value == (pos.value + 1)
                    && path.find(pos.surrounding[LEFT]) == path.end()) {

                auto new_path = path;
                queue.push({records[pos.surrounding[LEFT]], new_path}); //Push the element in the queue
            }
            if (is_valid(pos.surrounding[RIGHT], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[RIGHT]].value == (pos.value + 1)
                    && path.find(pos.surrounding[RIGHT]) == path.end()) {

                auto new_path = path;
                queue.push({records[pos.surrounding[RIGHT]], new_path}); //Push the element in the queue
            }
            if (is_valid(pos.surrounding[UP], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[UP]].value == (pos.value + 1)
                    && path.find(pos.surrounding[UP]) == path.end()) {

                auto new_path = path;
                queue.push({records[pos.surrounding[UP]], new_path}); //Push the element in the queue
            }
            if (is_valid(pos.surrounding[DOWN], max_col_idx, max_row_idx) 
                    && records[pos.surrounding[DOWN]].value == (pos.value + 1)
                    && path.find(pos.surrounding[DOWN]) == path.end()) {

                auto new_path = path;
                queue.push({records[pos.surrounding[DOWN]], new_path}); //Push the element in the queue
            }

        }

        if (flag) {
            ans2 += count;
        }

    }

    std::cout << "Answer 2: " << ans2 << std::endl;


    return 0;
}
