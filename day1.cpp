#include <stdio.h>
#include <stdlib.h>
#include <unordered_set>

void part1(char* filename)
{
	printf("*******\npart1\n*******\n");
	int   target = 2020;
	FILE* input  = fopen(filename, "r");
	std::unordered_set<int> nums;

	char str[5];
	while (fscanf(input, "%s\n", str) == 1)
	{
		int num = std::atoi(str);
		#ifdef DEBUG
		printf("%i\n", num);
		#endif
		nums.insert(num);
		if (nums.find(target - num) != nums.end())
		{
			printf("\nsolution: (%i, %i)\n", num, (target - num));
			printf("product: %i\n", num * (target - num));
			break;
		}
	}

	printf("\n");
	fclose(input);
}

void part2(char* filename)
{
	printf("*******\npart2\n*******\n");
	int   target = 2020;
	FILE* input  = fopen(filename, "r");
	std::unordered_set<int> nums;

	char str[5];
	while (fscanf(input, "%s\n", str) == 1)
	{
		int num = std::atoi(str);
		nums.insert(num);
	}
	fclose(input);

	for (const auto& n1 : nums)
	{
		for (const auto& n2 : nums)
		{
			if (n2 != n1 && nums.find(target - n2 - n1) != nums.end())
			{
				int n3 = target - n1 - n2;
				printf("\nsolution: (%i, %i, %i)\n", n1, n2,n3);
				printf("product: %i\n", n1 * n2 * n3);
				return;
			}
		}
	}
}

int main(int argc, char* argv[])
{
	part1(argv[1]);
	part2(argv[1]);
}
