# Go GRPC API (POC for Translation Service)


This POC is designed for a translation service.

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.15 or above**.

[GNU MAKE] is a program that automates the running of shell commands and helps with repetitive tasks. [the instructions](https://www.gnu.org/software/make/) to install it.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.

After installing Go and Docker, run the following commands to start experiencing this starter kit:

```shell
# download the POC
git clone https://github.com/jonayrodriguez/translation-service.git

cd translation-service


# start a MySQL database server and the service
make full-deploy

**NOTE:** if the service gives an error about connection to the DB. That is because it started too early. Please wait a seconds and run the this cmd: make run


```

At this time, you have the translation service running at `http://0.0.0.0:50051`. It provides the following endpoints:

* `GetTranslation`: Retrieve all the tranlated message by language, scope and key pattern.
* `AddTranslation`: Add a new translated message for a language, scope and key.
* `UpdateTranslation`: Update the translated message for the language, scope and key.
* `DeleteTranslation`: Delete a translated message for a language, scope and key. 


If you have `BloomRPC` or some gRPC client tools (e.g. [BloomRPC](https://github.com/uw-labs/bloomrpc), you can import the proto file to make the calls.

## Database Model
```
// Language table contains all the available languages
type Language struct {
	gorm.Model
	Name string `gorm:"unique;not null;size:2"`
	IETF string `gorm:"unique;not null;size:5"`
}

// Translation table contains all the transalation for the available languages
type Translation struct {
	gorm.Model
	LanguageName string   `gorm:"uniqueIndex:idx_translation"`
	Language     Language `gorm:"references:Name"`
	Scope        string   `gorm:"uniqueIndex:idx_translation;size:50"`
	Key          string   `gorm:"uniqueIndex:idx_translation;size:255"`
	Message      string   `gorm:"size:255;not null"`
}
```

## Project Layout

**TODO**

## Common Development Tasks

This section describes some common development tasks using this starter kit.

### Implementing a New Feature

**TODO**

### Working with DB Transactions

**TODO**