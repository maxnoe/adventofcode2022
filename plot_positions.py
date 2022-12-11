import numpy as np
import matplotlib.pyplot as plt

positions = np.genfromtxt("build/data.csv", delimiter=',')
x = positions[:, ::2]
y = positions[:, 1::2]
n_knots  = len(x[0])

cmap = plt.get_cmap("inferno")
colors = cmap(np.linspace(0.2, 1.0, n_knots))


fig = plt.figure(facecolor='k', figsize=(25.6, 14.4), dpi=100)
ax = fig.add_axes([0, 0, 1, 1], facecolor='k')
ax.set_axis_off()
ax.set_aspect(1)

for i, color in enumerate(colors):
    ax.plot(x[:, i], y[:, i], color=color)

fig.savefig('/home/maxnoe/rope.png')
