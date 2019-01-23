<p align="center"><a href="https://www.shyft.network/">
	<img src="https://www.shyft.network/images/shyft-logo-horizontal-tm.svg" alt="Shyft" height="100px" align="center">
</a></p>

# Shyft Block Explorer UI

This repository contains the Shyft Block Explorer UI which is primarily used to display blockchain information from Shyft's go-empyrean blockchain. It is served by an API that is communicating with a postgres instance on the go-empyrean chain.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

It is required to have Shyft's go-empyrean blockchain postgres instance running and Shyft's Block Explorer API in order to utilize the UI. You can run the API by following the instructions [here](https://github.com/ShyftNetwork/blockexplorer_api) and Shyft's go-empyrean blockchain by following the instructions [here](https://github.com/ShyftNetwork/go-empyrean).

In addition you will need `node 8.9.4` or greater.

```
1. go-empyrean postgres instance
2. Shyft Block Explorer API running
3. > node 8.9.4
```

### Installing

To get a development environment running locally please follow the below steps:

```
1. git clone git@github.com:ShyftNetwork/blockexplorer_ui.git
2. npm install
3. npm run start
```

**You should see the development server starting and listening on port 3000!**

## Running the tests

To run the tests run the following command:

```
> npm run tests
```

