#include<stdio.h>
#include<math.h>
#include<string.h>
//用迭代法求x=根号a。求平方根的迭代公式 
int main()
{
	float a,x0,x1;
	printf("please input a:");
	scanf("%f",&a);//a需要人工输入
	x0=a/2; //给x0赋初始值，也可以选择其他值
	x1=(x0+a/x0)/2;
	do{
		x0=x1;
		x1=(x0+a/x0)/2;
	}while(fabs(x1-x0)>=1e-5);
	printf("The square root of %5.2f is %8.5f\n",a,x1);
	return 0;
}
