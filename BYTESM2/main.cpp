#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

void solve(int h, int w, vector< vector<int> >& values)
{
    int hh = h + 1;
    int ww = w + 2;
    vector< vector<int> > dp;
    vector<int> row;
    for(int x=0;x<ww;x++) {
        row.push_back(0);
    }
    for(int y=0;y<hh;y++) {
        dp.push_back(row);
    }

    for(int y=1;y<hh;y++) {
        for(int x=1;x<ww-1;x++) {
            dp[y][x] = values[y-1][x-1] + max( max(dp[y-1][x-1], dp[y-1][x]), dp[y-1][x+1] );
        }
    }

    cout << *max_element(dp[hh-1].begin(), dp[hh-1].end()) << endl;
}

int main()
{
    int numTests = 0;
    cin >> numTests;
    for(int i=0;i<numTests;i++) {
        int h, w = 0;
        cin >> h >> w;

        vector< vector<int> > values;
        for(int y=0;y<h;y++) {
            vector<int> row;
            for(int x=0;x<w;x++) {
                int v = 0;
                cin >> v;
                row.push_back(v);
            }
            values.push_back(row);
        }

        solve(h, w, values);
    }
}