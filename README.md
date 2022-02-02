currently work in progress and incomplete

screenshot
=============

screenshot capturing and uploading program

features
-------------
- written in go so easy to install with zero dependency fuss
- simple gallery generator which makes thumbnails, templates for themes
- upload to s3 API

Usage
-------------
        NAME:
        screenshot - screenshot

        USAGE:
        screenshot [global options] command [command options] [arguments...]

        COMMANDS:
        config   show screenshot config
        capture  capture a screenshot
        gallery  make a gallery of images
        help, h  Shows a list of commands or help for one command

        GLOBAL OPTIONS:
        --help, -h  show help (default: false)
   
configuration example
---------------------------------

        datadir: /tmp/shots
        capture:
            command: import
        upload:
            interface: s3
            access_key: [aws access key]
            secret_key: [aws secret key]
            bucket: [ aws bucket ]
