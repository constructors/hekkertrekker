Hekker Trekker
==============

Automates our Mercurial and Pivotal Tracker integration.

All it needs are two config files.

Configuration
-------------

There's `~/.hekkertrekker`:

    {"token":"yoursecrettoken",
     "newbranchcommitmsg":"Accepted story %d.",
     "delivercommitmsg":"Delivered story %d.",
     "donecommitmsg":"Done story %d.",
     "donelabel":"live",
     "name":"Your Name Exactly As Pivotal Tracker Knows It"}

And in your source dir, also called `.hekkertrekker`:

    {"projectid":12345,"stagingbranch":"staging"}

Usage
-----

* `ht accept` - Accept a new story and create a new branch.
* `ht deliver` - Deliver current branch and merges to staging.
* `ht done` - Adds a tag to the story and merges to default.

Installation
------------

Make sure you've got go installed. Run `make`. Copy the resulting `ht` binary
somewhere into your `$PATH`.

Tests
-----

Are you kidding me? This is a hack.
