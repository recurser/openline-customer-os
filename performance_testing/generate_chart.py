import matplotlib.pyplot as plt
import pandas as pd
import numpy as np

# Load data from CSV file
data = pd.read_csv('average_response_times.csv')

# Extract x-axis values from the first column
x = data.iloc[:, 0]

# Extract y-axis values from all subsequent columns
y = data.iloc[:, 1:]

# Generate the chart
for i, column in enumerate(y.columns):
    x = np.arange(len(y[column]))
    plt.plot(x, y[column], marker='o', label=column)

# Set x-axis tick labels to the values from the first column
plt.xticks(np.arange(len(data)), data.iloc[:, 0])

# Add legend to show the lines for each column
plt.legend()

# Set x-axis tick labels to the values from the first column, aligned vertically
plt.xticks(range(len(x)), x, rotation='vertical')

# Save the chart as an image
plt.savefig('average_response_times.png')