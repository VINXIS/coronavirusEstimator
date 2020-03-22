# coronavirusEstimator
This code uses the method [explained thoroughly by 3Blue1Brown in his video regarding Exponential Growth and epidemics](https://www.youtube.com/watch?v=Kas0tIxDvrg)

The formula used for this is next day total cases = previous total * (1 + average number of people someone infected is exposed to * (1 - previous total / population))

These results assume that no extra measures will be taken by the population from now on as the calculations assume the "average number of people someone infected is exposed to" value to be constant.

Some dummy data is provided based on around the afternoon of March 21 2020 UTC-7 to show how to format the data to use for testing. Let me know how I can improve the estimated total cases!