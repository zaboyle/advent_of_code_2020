# designate which compiler to use
CXX			 = g++

#Default Flags
CXXFLAGS     = -std=c++14 -pthread # -Wconversion -Wall -Werror -Wextra -pedantic

# make release - will compile "all" with $(CXXFLAGS) and the -O3 flag
#				 also defines NDEBUG so that asserts will not check
release: CXXFLAGS += -O3 -DNDEBUG
release: all

# make debug - will compile "all" with $(CXXFLAGS) and the -g flag
#              also defines DEBUG so that "#ifdef DEBUG /*...*/ #endif" works
debug: CXXFLAGS += -g3 -DDEBUG
debug: clean all

all: day1 day2 day3

day1:
	$(CXX) $(CXXFLAGS) day1.cpp -o day1

day2:
	go build -o day2 day2.go

day3:
	go build -o day3 day3.go

clean:
	rm -rf day1 day2 day3