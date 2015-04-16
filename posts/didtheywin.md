title: DidTheyWin README
date: 1-28-2015
<-->
DidTheyWin?
==========

A one page Sinatra app to see if an NBA team won their last game.
Created by Bryce Kalow, http://brycekalow.name
Version 1.0

Usage
-----

First create a file named `app.yml` in `/config` containing:

```
keys:
  xmlstats: [YOUR XMLSTATS KEY HERE]
```

Next, run:
`bundle install`

Finally:
`bundle exec shotgun`
Or:
`rackup`

Testing
-------

To run the test file, from the root of the app folder:
`rspec spec/app_spec.rb`

Credits
-------

Data provided by the xmlstats API: http://erikberg.com/api
