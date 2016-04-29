# BeenVerified API Challenge

The following code is an API that delivers songs and genres data from an SQLite database that can be retrieved [here](https://s3.amazonaws.com/bv-challenge/jrdd.db).

## Getting Started

These steps will help you with the installation process to set up the server locally in your machine.

### Prerequisities

You need to make sure you have installed the *Go Programming Language* 1.6.X in your system. If not, you can download and get help installing it [here](https://golang.org/doc/install).

You can check the version installed in your system with the following command:
```
go version
```


You will also need *glide* installed and you can do it following the installation instructions in its [repository](https://github.com/Masterminds/glide).

You can verify the version of glide using:
```
glide -version
```


### Installing

#### Cloning the repository
You can clone this repository to create a local copy on your computer with:
```
git clone https://github.com/brunosarm70/beenverified_challenge $GOPATH/src/github.com/brunosarm70/beenverified-challenge
```

#### Installing the dependencies

You will have to install all the neccessary dependencies to run the project, using the glide.yaml file.
Make sure you are in the folder where glide.yaml is located.
```
cd $GOPATH/src/github.com/brunosarm70/beenverified-challenge
```
And then install the dependencies:
```
glide install
```

#### Compiling with the *go* tool

The following command will create an executable file named *beenverified-challenge* in the bin folder of your workspace:
```
go install github.com/brunosarm70/beenverified-challenge
```

Execute the command to run the server:
```
$GOPATH/bin/beenverified-challenge
```
Now the server will be listening on 127.0.0.1 port 8000.


## Authors

* **Bruno Sarmiento**  


## Acknowledgments

* 
