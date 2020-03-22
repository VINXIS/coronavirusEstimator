import glob
import json
import matplotlib.pyplot as plt
import numpy as np
import os
import pandas as pd
import re

fig = plt.figure(figsize=[36, 24])

files = glob.glob('./results/*.json')

for file in files:
    with open(file, 'r') as f:
        data = json.load(f)
    dataframe = pd.DataFrame(data).iloc[1:]
    plt.plot(dataframe.index, dataframe, label=re.sub('.json', '', os.path.basename(file)))

plt.xticks(np.arange(0, len(dataframe.index), 1))
plt.grid(b=True, which='both')
plt.legend()

plt.savefig('./images/results.png', bbox_inches='tight')

plt.yscale('log')
plt.savefig('./images/resultsLog.png', bbox_inches='tight')