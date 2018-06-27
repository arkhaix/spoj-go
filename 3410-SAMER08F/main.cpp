#include <iostream>
using namespace std;

int main()
{
    int dp[101];
	dp[0]=0;
	for(int i=1;i<=100;i++) {
		dp[i]=dp[i-1]+(i*i);
	}
	
	while(true)
	{
		int n;
		cin >> n;
		if(n == 0) break;
		cout << dp[n] << endl;
	}
	// your code here

	return 0;
}