# coding: utf-8
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation

positions = np.genfromtxt("data.csv", delimiter=',')
x = positions[:, ::2]
y = positions[:, 1::2]

print(x)

fig = plt.figure()
fig.show()

ax = fig.add_axes([0, 0, 1, 1])
ax.set_aspect(1)
ax.set_xlim(x.min() - 0.5, x.max() + 0.5)
ax.set_ylim(y.min() - 0.5, y.max() + 0.5)
ax.set_axis_off()

line, = ax.plot(x[0], y[0], marker='o', ls=':')


def update(row): 
    print(ax.get_xlim(), ax.get_ylim())
    print(x[row, 0], y[row, 0])
    line.set_xdata(x[row])
    line.set_ydata(y[row])
    return line,


ani =  FuncAnimation(fig, update, frames=len(x), interval=100, blit=True)
ani.save("correct.mp4")
