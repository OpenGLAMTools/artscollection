# artscollection
A flexbile way to store and read data

## Configuration

### Program

The configuration is all inside the `.artscollection.yaml` file.

The structure of the configuration looks like:

    author: authorGithub
    serverPort: :8081
    collections:
            myCollection: path/to/collection
            anotherCollection: path/to/collection

### Collection

Inside of every collection there are also individual configuration possibilities. In the root folder of the collection a `conf.yaml` file is needed. This file defines all the default fields for the collection.

If there is a field with the key `title` that value is used inside the navigation. Otherwise the folder name is used as title.

    -  key: title
        name: Name
        type: string
        render: 
        group: a
        order: 10
    -  key: desc
        name: Description
        type: string
        render: textarea
        group: a
        order: 20 

Following types are allowed:

* string
* int
* bool
* list

The `string` type supports the `render` property. If `render: textarea` the input field is rendered as a textarea.

