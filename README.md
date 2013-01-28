Hekker Trekker
==============

Automates our Mercurial and Pivotal Tracker integration.

All it needs is two config files.

Configuration
-------------

There's `~/.hekkertrekker`:

    {"token":"yoursecrettoken",
     "newbranchcommitmsg":"Accepted story %d.",
     "delivercommitmsg":"Delivered story %d.",
     "closecommitmsg":"Closing branch.",
     "donecommitmsg":"Done story %d.",
     "donelabel":"live",
     "name":"Your Name Exactly As Pivotal Tracker Knows It"}

And in your source dir, also called `.hekkertrekker`:

    {"projectid":12345,"stagingbranch":"staging"}

Usage
-----

* `ht start` - Starts a new story and creates a new branch.
* `ht deliver` - Delivers current branch and merges to staging.
* `ht done` - Adds a tag to the story and merges to default.

Installation
------------

Make sure you've got go installed. Run `make`. Copy the resulting `ht` binary
somewhere into your `$PATH`.

Tests
-----

Are you kidding me? This is a hack.
