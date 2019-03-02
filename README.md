# infest

Provides an API for the Baltimore City Recent Food Establishment Closures page
[here](http://health.baltimorecity.gov/environmental-health/recent-food-establishment-closures)

The application scrapes data from the page periodically and exposes that information in a queryable REST api.

As of early 2017, this particular data is not exposed on the Open Baltimore Portal, and having this data easily accessible
will enable a wider variety of civic hacking the future. Examples could include notification systems, search engines/food review
integrations, and geographic visualizations

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

## running

The below environment variables must be set for the app to run. Defaults are given as well.

* `CLOSURES_URL` - the url of the health department page with the restaurant closures
* `CLOSURES_SCHEDULE` - a schedule compatible with the [robfig/cron](https://github.com/robfig/cron) library. Default is every minute.
* `PORT` - the port that the application will run on
* `GO_ENV` - Either "development" or "production" depending on the environment. Default is development if unset.

## dependencies

The full list is in go.mod. In development and test the app uses sqlite to keep it lightweight. In production, postgres is used.
