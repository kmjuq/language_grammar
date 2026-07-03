# Makefile for c/ project
# 编译 c/main.c -> c/main

CC      ?= cc
CFLAGS  ?= -std=c11 -Wall -Wextra -O2
LDFLAGS ?=
LDLIBS  ?= -lm

SRC_DIR := .
TARGET  := main
SRC     := $(SRC_DIR)/main.c
OBJ     := $(SRC:.c=.o)

.PHONY: all run clean rebuild

all: $(TARGET)

$(TARGET): $(OBJ)
	$(CC) $(CFLAGS) $(LDFLAGS) -o $@ $^ $(LDLIBS)

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

run: $(TARGET)
	./$(TARGET)

clean:
	rm -f $(OBJ) $(TARGET)

rebuild: clean all
