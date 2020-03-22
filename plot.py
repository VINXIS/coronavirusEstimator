import glob
import json
import matplotlib.pyplot as plt
import numpy as np
import os
import pandas as pd
import re

fig, axarr = plt.subplots(3, sharex=True, figsize=[36, 36])

with open('data.json', 'r') as f:
    knownData = json.load(f)

files = glob.glob('./results/*.json')
for file in files:
    with open(file, 'r') as f:
        data = json.load(f)
    dataframe = pd.DataFrame(data).iloc[1:]
    countryName = re.sub('.json', '', os.path.basename(file))
    for country in knownData:
        if country['Name'] == countryName:
            countryPopulation = country['Population']
            break
            
    axarr[0].plot(dataframe.index, dataframe, label=countryName)
    axarr[1].plot(dataframe.index, dataframe, label=countryName)
    axarr[2].plot(dataframe.index, dataframe / countryPopulation, label=countryName)

plt.title('Expected total cases of COVID-19 per days from now')

axarr[0].set_xlabel('days from now')
axarr[0].set_ylabel('# of people infected')

axarr[1].set_xlabel('days from now')
axarr[1].set_ylabel('# of people infected')
axarr[1].set_yscale('log')

axarr[2].set_xlabel('days from now')
axarr[2].set_ylabel('%\ of people infected')

plt.xticks(np.arange(0, len(dataframe.index), 1))
plt.grid(b=True, which='both')
plt.legend()

plt.savefig('./results/results.png', bbox_inches='tight')