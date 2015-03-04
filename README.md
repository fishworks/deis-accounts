# Deis Accounts

deis-accounts is a plugin for the Deis Client to help manage multiple accounts.

## Installation

    $ make && sudo make install

## Usage

To add accounts:

    $ deis accounts:add
    login URL: http://local3.deisapp.com
    username: bacongobbler
    password: ******

To switch contexts to a different account:

    $ deis accounts:set bacongobbler

To list accounts:

    $ deis accounts:list
    bacongobbler
    testuser

To remove an account:

    $ deis accounts:remove testuser
    Account removed.
