#include <iostream>
#include <vector>
using namespace std;

struct node {
  int v;
  node *l, *r;
  int li, ri;
};

vector<int> a;
node* build(int i, int j) {
  node* n = new node();
  n->v = 0;
  n->li = i;
  n->ri = j;
  n->v = a[i];
  if(j-i > 0) {
    int mid = (i+j)/2;
    n->l = build(i, mid);
    n->r = build(mid+1, j);
    n->v = max(n->l->v+n->r->v, max(n->l->v, n->r->v));
  }
  return n;
}

int solve(int x, int y, node* n) {
  if(y<n->li || x>n->ri) {
    return -1;
  }
  if(n->li >= x && n->ri <= y) {
    return n->v;
  }
  return max(solve(x,y,n->l), solve(x,y,n->r));
}

int main() {
  int n;
  cin >> n;
  a.resize(n+1);
  a[0]=0;
  for(int i=0;i<n;i++) {
    cin >> a[i+1];
  }
  node* tree = build(0, n);
  int numQueries;
  cin >> numQueries;
  for(int i=0;i<numQueries;i++) {
    int x, y;
    cin >> x >> y;
    cout << solve(x, y, tree) << endl;
  }
  return 0;
}