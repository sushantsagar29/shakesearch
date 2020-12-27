# ShakeSearch

Welcome to the Pulley Shakesearch Take-home Challenge! In this repository,
you'll find a simple web app that allows a user to search for a text string in
the complete works of Shakespeare.

You can see a live version of the app at
https://pulley-shakesearch.herokuapp.com/. Try searching for "Hamlet" to display
a set of results.

In it's current state, however, the app is just a rough prototype. The search is
case sensitive, the results are difficult to read, and the search is limited to
exact matches.

## Your Mission

Improve the search backend. Think about the problem from the **user's perspective**
and prioritize your changes according to what you think is most useful. 

## Evaluation

We will be primarily evaluating based on how well the search works for users. A search result with a lot of features (i.e. multi-words and mis-spellings handled), but with results that are hard to read would not be a strong submission. 


## Submission

1. Fork this repository and send us a link to your fork after pushing your changes. 
2. Heroku hosting - The project includes a Heroku Procfile and, in its
current state, can be deployed easily on Heroku's free tier.
3. In your submission, share with us what changes you made and how you would prioritize changes if you had more time.

---
---

## Changes done 

#### [Application url](https://shakesearch-sushant.herokuapp.com/)

### Explanation
1. `Aa` symbol with checkbox on the ui is for case sensitive search
2. `W` symbol with checkbox on ui is for searching complete word

### Functional
1. Highlighted text which is being searched since reading through the result was difficult otherwise
2. Added option for user to toggle for case sensitive search
3. Added option for user to toggle for searched text to be complete word or part of another word
4. Added loader on ui, other wise scenarios where search result was huge used to take time to render and made the ui unresponsive
5. Added search count result, and made the results load in tabular form making it easier to read
6. Fixed existing bug in the code where if the queried word did not have 250 characters before or after it, was breaking the code for slice bounds out of range  
### Non functional
1. Displayed error messages in non successful fetch scenarios
2. Refactored code to have separate files with separate responsibilities (handler, service, model)


### Further improvements I had planned if given more time
1. Added pagination, since queries with thousands of results are taking time to load
2. Would have parallelized the search in backend by breaking the entire data into chunks
3. Moved UI to reactJs where handling of pagination would have become simple, handling of loader which is currently being done in a very vanilla approach would have became simple by using state, even the rendering of data in table would would have been through a component instead of appending into string 
4. Test cases. First priority since the functionalities are increasing and points of failure would increase if not tested properly
5. Moved all error messages and other constants to a different constant file keeping the code clean of plain strings
6. Would have added logging in real world scenario
7. Add debounce to search operation and remove the submit button (provides better ux in some cases)
