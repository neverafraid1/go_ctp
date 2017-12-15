MAKE = make
CXX = g++
CFLAGS = -I./include -I./ -I../
CFLAGS += -O0
CXXFLAGS = -Wall -fpermissive
CXXFLAGS += -fPIC
CXXFLAGS += -rdynamic
CXXFLAGS += -g


LIBLDFLAGS = -shared
LDFLAGS = -L/usr/local/lib -lthostmduserapi -lthosttraderapi

.PHONY: all install clean cleantest coverage html

SRC = $(wildcard *.cpp)

OBJS = $(notdir $(SRC:.cpp=.o))

TARGET := libctpwrapper.so

DEPEND_LIBS = ./lib64/libthosttraderapi.so ./lib64/libthostmduserapi.so

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CXX) $(CFLAGS) $(CXXFLAGS) -o $@ $^ $(LIBLDFLAGS) $(LDFLAGS)

%.o: %.cpp
	$(CXX) $(CFLAGS) $(CXXFLAGS) -c $<

%.o: %.cc
	$(CXX) $(CFLAGS) $(CXXFLAGS) -c $<

clean:
	-rm -f  *.o $(TARGET)
