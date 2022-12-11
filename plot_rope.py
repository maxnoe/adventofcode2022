# coding: utf-8
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
from tqdm import tqdm

positions = np.genfromtxt("build/data.csv", delimiter=',')
x = positions[:, ::2]
y = positions[:, 1::2]

print(x)

fig = plt.figure(figsize=(12.8, 7.2), dpi=100, facecolor='k')
fig.show()

ax = fig.add_axes([0, 0, 1, 1], facecolor='k')
ax.set_aspect(1)
ax.set_xlim(x.min() - 0.5, x.max() + 0.5)
ax.set_ylim(y.min() - 0.5, y.max() + 0.5)
ax.set_axis_off()


cmap = plt.get_cmap('inferno')
c = np.linspace(0.2, 1.0, len(x[0]))
colors = cmap(c)

line, = ax.plot(x[0], y[0], ls=':', marker='none', color='k')
lines = [
    ax.plot([], [], color=colors[i], marker='.', ms=3, alpha=0.2)[0]
    for i in range(len(x[0]))
]
points = ax.scatter(x[0], y[0], c=c, cmap=cmap, s=50, zorder=3, vmin=0, vmax=1)

bar = tqdm(total=len(x) + 1)

def update(row): 
    bar.update(1)
    line.set_xdata(x[row])
    line.set_ydata(y[row])
    points.set_offsets(np.column_stack([x[row], y[row]]))
    for i, l in enumerate(lines):
        l.set_data(x[:row + 1, i], y[:row + 1, i])
    return line, points, *lines


fps = 100
interval = 1000 / fps
ani =  FuncAnimation(fig, update, frames=len(x), interval=interval, blit=True)
ani.save("build/movie.mp4")
