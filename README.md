currently work in progress and incomplete

screenshot
=============

screenshot capturing and uploading program

features
-------------
- written in go so easy to install with zero dependency fuss
- simple gallery generator which makes thumbnails
- templates for themes

configuration
---------------------------------

        datadir: /tmp/shots
        capture:
            command: import
        upload:
            interface: s3
            access_key: [aws access key]
            secret_key: [aws secret key]
            bucket: [ aws bucket ]
