from matplotlib import pyplot as plt
import numpy as np
import sys
import re

nodes = []
best_path = []

def read_file(filepath):
    with open(filepath, 'r', encoding='utf-8') as f:
        input = " "
        while input:
            input = f.readline()
            input = re.sub(r'[\[\]]', '', input)
            input = input.split()
            if len(input) > 0:
                nodes.append(input)


def plotting_problem(nodes):
    x = [int(node[0]) for node in nodes]
    y = [int(node[1]) for node in nodes]    
    plt.scatter(x, y, s=25, alpha=0.75)
    plt.title('Problem Nodes')
    plt.xlabel('X')
    plt.ylabel('Y')
    plt.show()

#TBD
def plotting_problem_with_solution(nodes, best_path = None):
    colors = ['b', 'g', 'r', 'c', 'm', 'y', 'k', 'w']
    color_index = 0
    i = 1
    for i in range(len(path)):
        plt.plot([int(params[path[i - 1]]['x']), int(params[path[i]]['x'])], [int(params[path[i - 1]]['y']),
                                                                              int(params[path[i]]['y'])], 'o', ls='-',
                 color=colors[color_index], alpha=0.5, ms=4)
        if len(path) > i > 0 == path[i]:
            color_index += 1
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
