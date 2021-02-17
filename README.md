# M-BDX-002
Epitech 5th year go module

# Installation

If you do not have Golang installed, go to https://golang.org/doc/install?download and find the adequate version.

Following the installation a "go" folder will be created at the root of your user session

    example:
      - on mac: /Users/<username/

In this file create if not existing a "github.com" folder and then a "LaurentColoma" folder.

You can then clone the project in the last folder created.

Jump into the clone folder and you can either do:

    go run main.go examples/<example_file_provided>
Or

    go build
    ./M-BDX-002 examples/<example_file_provided>
    
In the case the above command return "go: command not found" do the following:

    export PATH=/usr/local/go/bin:$PATH

Retry and it should be working as expected, and in the case you encounter an other error try "man google"


# Package

Folders:
  - examples: Custom examples to test the project
  - gameData: Data structure and movement related method
  - gameLoop: The heart of the program, the main feature of the games are done here
  - parsing: Handle the parsing
  - pathFinding: Handle the movement of the pallet in the warehouse, it find the most optimized path.

Files:
  - M-BDX-002: binary
  - main.go: Start of every go program
  - projet de gestion d'entrpôt.pdf: Subject
  - README.md: The file you are currently reading explaning EVERYTHING

# Strategy

*Simple is better than complex*

We look for the lightest's Parcel, we then make the PalleTruck head to the Parcel using our pathfinding and we pick it.

We then head to the Truck position, and on arrival we drop it off or we wait for the Truck to be available.

And we go again until it is over.


By Benjamin VIGUIER, Arnaud CLERC & Laurent COLOMA
