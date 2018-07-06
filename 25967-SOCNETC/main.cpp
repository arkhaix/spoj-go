#include <iostream>
#include <vector>
using namespace std;

vector<int> parent;
vector<int> sizes;
int maxSize = 0;

int root(int x) {
    while(parent[x] != x) {
        x = parent[x];
    }
    return x;
}

void merge(int x, int y) {
    int rx = root(x);
    int ry = root(y);
    if(sizes[rx]+sizes[ry] <= maxSize) {
        if(sizes[rx] < sizes[ry]) {
            parent[rx] = ry;
            sizes[ry] += sizes[rx];
        } else {
            parent[ry] = rx;
            sizes[rx] += sizes[ry];
        }
    }
}

void query(int x, int y) {
    if(root(x) == root(y)) {
        cout << "Yes" << endl;
    } else {
        cout << "No" << endl;
    }
}

void printSize(int x) {
    cout << sizes[root(x)] << endl;
}

int main()
{
    int numUsers = 0;
    cin >> numUsers;
    numUsers++;
    cin >> maxSize;

    parent.resize(numUsers, 0);
    sizes.resize(numUsers, 1);

    for(int i=0;i<numUsers;i++) {
        parent[i] = i;
    }

    int numQueries = 0;
    cin >> numQueries;

    for(int i=0;i<numQueries;i++) {
        char command = 0;
        cin >> command;

        if(command == 'A') {
            int x, y;
            cin >> x >> y;
            merge(x, y);
        } else if(command == 'E') {
            int x, y;
            cin >> x >> y;
            query(x, y);
        } else /* S */ {
            int x;
            cin >> x;
            printSize(x);
        }
    }

    return 0;
}