from matplotlib import pyplot as plt
import numpy as np
import sys
import re

nodes = []
best_path = []


def read_file(filepath):
    global best_path
    with open(filepath, 'r', encoding='utf-8') as f:
        input = " "
        indices = []
        while input:
            input = f.readline()
            input = re.sub(r'[\[\]]', '', input)
            input = input.split()
            if len(input) > 0:
                nodes.append(input)
            else:
                indices = f.read()
        if len(indices) > 0:
            indices = indices.split('\n')
            indices.pop()
    best_path = [int(i) for i in indices]


def plotting_problem(nodes):
    x = [int(node[0]) for node in nodes]
    y = [int(node[1]) for node in nodes]
    plt.scatter(x, y, s=25, alpha=0.75)
    plt.title('Problem Nodes')
    plt.xlabel('X')
    plt.ylabel('Y')
    plt.show()


def plotting_problem_with_solution(nodes, path=None):
    color_index = 0
    i = 1
    for i in range(len(path)):
        plt.plot([int(nodes[path[i - 1]][0]), int(nodes[path[i]][0])],
                 [int(nodes[path[i - 1]][1]), int(nodes[path[i]][1])],
                 'o', ls='-', alpha=0.5, ms=4)
        # if len(path) > i > 0 == path[i]:
            # color_index += 1
    plt.title('TSP Problem Result')
    plt.xlabel('X')
    plt.ylabel('Y')
    plt.show()


if __name__ == "__main__":
    try:
        read_file(sys.argv[1])
        if len(best_path) > 0:
            plotting_problem_with_solution(nodes, best_path)
        else:
            plotting_problem(nodes)
    except:
        print("Filepath error!")
