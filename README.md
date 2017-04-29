# infest

Provides an API and email notifications for the Baltimore City Recent Food Establishment Closures page 
[here](http://health.baltimorecity.gov/environmental-health/recent-food-establishment-closures)

The application scrapes data from the page periodically and exposes that information in a queryable REST api.

## using

The app is deployed [here](https://balto-restaurant-closures.herokuapp.com/)

The valid url params are as follows:

* per_page - the number of results returned per page - default is 20
* page - the page offset - default is 1
* sort - accepts lowercase param names followed by either ".asc" or ".desc" (e.g. "closuredate.desc")
* name - filters for an exact match on the name field
* startDate - shows all records for which closuredate is after the given date. Format is m/d/yyyy
* endDate - shows all records for which closuredate is before the given date. Format is m/d/yyy
* reason - does a full text search in the reason field for the given item