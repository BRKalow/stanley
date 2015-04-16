title: Swish README
date: 1-26-2015
<-->
Swish
=====

Swish is a code snippet sharing application. It was inspired by the concept of [Dribbble](https://dribbble.com), with the focus on developers instead of designers. Swish is also my final project for the INET 3350 (Ruby/Rails) class at the University of Minnesota taught by John Norman. This was a purely educational project and has no real intended longevity, or promises of quality.

Technologies
------------

Swish uses Ruby on Rails 4.1, with Postgres configured as its defaulted database driver. A collection of gems are also used in the project, which can be founded listed in the Gemfile.

Local Installation
------------

To install swish, clone the repository:

```
git glone git@github.com:BRKalow/swish.git && cd swish
```

Then run bundle install:

```
bundle install
```

Setup the database:

```
rake db:create && rake db:migrate
```

Finally run the development server:

```
bundle exec rails s
```

Navigating to `localhost:3000` should open the application.

Heroku Deployment Link
----------------------

Check it out [here](http://swish.brycekalow.name).
