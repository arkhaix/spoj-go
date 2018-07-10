#include <iostream>
#include <vector>
using namespace std;

struct FenwickTree
{
    vector<long long> elements;

    void buildFrom(vector<int>& arr)
    {
        elements.resize(arr.size(), 0LL);
        for(size_t i = 0; i < arr.size(); i++) {
            update(i, arr[i]);
        }
    }

    void update(int index, int value)
    {
        size_t n = elements.size();
        for(size_t i = index; i < n; i = i | (i+1)) {
            elements[i] += value;
        }
    }

    long long query(int left, int right)
    {
        return sum(right) - sum(left-1);
    }

    long long sum(int index)
    {
        long long res = 0LL;
        while(index > 0)
        {
            res += elements[index];
            index = index & (index + 1);
            index--;
        }
        return res;
    }
};

int main()
{
   int numElements;
   cin >> numElements;

   vector<int> arr;
   arr.push_back(0LL);
   int n = 0;
   for(int i = 1; i <= numElements; i++) {
       cin >> n;
       arr.push_back(n);
   }

    FenwickTree tree;
    tree.buildFrom(arr);

   int numOperations;
   cin >> numOperations;

   char operation = 0;
   int a = 0;
   int b = 0;

    for(int i = 0; i < numOperations; i++) {
        cin >> operation;
        cin >> a;
        cin >> b;

        if(operation == 'q') {
            cout << tree.query(a, b) << endl;
        } else {
            tree.update(a, b);
        }
   }

    return 0;
}