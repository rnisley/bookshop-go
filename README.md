# bookshop-go

A simple book store API in need of input validation/sanitization.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:

- [Javascript](https://github.com/andey-robins/bookshop-js)
- [Rust](https://github.com/andey-robins/bookshop-rs)

## Versioning

`bookshop-go` is buit with:

- go version go1.19.3 darwin/arm64

## Usage

Start the api using `go run main.go`.

I recommend using [`httpie`](https://httpie.io) for testing of HTTP endpoints on the terminal. Tutorials are available elsewhere online, and you're free to use whatever tools you deem appropriate for testing your code.

## Analysis of Existing Code

New customers start with a default 5.0 as thier account balance.
Customers must know their ID to update shipping info as they can't look it up as part of the update and no api call exists to look up thier ID.
Customers must know their ID to look up account balance as they can't look it up as part of the action and no api call exists to look up thier ID.
If a customer orders multiple copies of the same book the db may behave unpredictably.
Database never created or connected to.

