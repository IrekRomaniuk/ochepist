
## Under Development 

#### Arguments

 ./ochepist_linux_386 -h

Copyright 2017 @IrekRomaniuk. All rights resgit brerved.

Usage of ./ochepist_linux_386:

  -c string

        comment added to objects (default "Created by ochepist with dbedit")

  -g string

        Object group to include retrieved IP addresses (default "g-ochepist-temp")

  -p string

        path to created dbedit file (default "./results/")

  -t string

        path to dbedit tempalte files (default "templates/*")

  -url string

        URL to pull IP addresses from (default "https://minemeld/feeds/office365_IPv4s")

  -v    Prints current version

#### Example

[Expert@provider1:0]# ./ochepist_linux_386 -url="https://minemeld/feeds/office365_IPv4s" -g="g-o365"

wrote 357258 bytes

[Expert@provider1:0]# head results/g-o365-dbedit.txt 

create network_object_group g-o365

create address_range r104.210.43.160-104.210.43.160

modify network_objects r104.210.43.160-104.210.43.160 ipaddr_first 104.210.43.160

modify network_objects r104.210.43.160-104.210.43.160 ipaddr_last 104.210.43.160

modify network_objects r104.210.43.160-104.210.43.160 comments "Created by ochepist with dbedit"

update network_objects r104.210.43.160-104.210.43.160

addelement network_objects g-o365 '' network_objects:r104.210.43.160-104.210.43.160

update network_objects g-o365

create address_range r104.41.155.129-104.41.155.129

modify network_objects r104.41.155.129-104.41.155.129 ipaddr_first 104.41.155.129

[Expert@provider1:0]# dbedit -local -globallock -f results/g-o365-dbedit.txt
