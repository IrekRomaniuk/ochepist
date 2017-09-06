
## Under Development 

### Arguments

```
 ./ochepist_linux_386 -h

Copyright 2017 @IrekRomaniuk. All rights reversed.

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
```

### Example

```
[Expert@provider1:0]# ./ochepist_linux_386 -url="https://minemeld/feeds/office365_IPv4s" -g="g-o365"
wrote 357258 bytes

[Expert@provider1:0]# head results/g-o365-dbedit.txt
delete network_objects g-o365
create network_object_group g-o365
delete network_objects r104.210.43.160-104.210.43.160
create address_range r104.210.43.160-104.210.43.160
modify network_objects r104.210.43.160-104.210.43.160 ipaddr_first 104.210.43.160
modify network_objects r104.210.43.160-104.210.43.160 ipaddr_last 104.210.43.160
modify network_objects r104.210.43.160-104.210.43.160 comments "Created by ochepist with dbedit"
update network_objects r104.210.43.160-104.210.43.160
update network_objects g-o365

[Expert@provider1:0]# tail results/g-o365-dbedit.txt
update_all
savedb

[Expert@provider1:0]# dbedit -local -ignore_script_failure -globallock -f results/g-o365-dbedit.txt
```
### ToDo

+ remove empty lines and spaces from end of lines
+ missing end of lines in g-ochecpist-test-dbedit.txt (including differences between Win and Linux)
+ return hash of url body and keep in env variable in order to skip if not changes
+ add #comments in template files i.e. do not add spaces or empty lines
+ review https://www.cpug.org/forums/showthread.php/21002-How-to-handle-Office365-IP-addresses
+ comply with https://support.content.office.net/en-us/static/O365IPAddresses.xml not only minemeld
+ Conversion from CIDR prefix notation to IP address/subnet mask (dot-decimal), see https://github.com/EvilSuperstars/go-cidrman/blob/master/ipv4.go npm netmask


### References

1. https://sc1.checkpoint.com/documents/R77/CP_R77_CLI_ReferenceGuide_WebAdmin/html_frameset.htm?topic=documents/R77/CP_R77_CLI_ReferenceGuide_WebAdmin/105997
2. https://supportcenter.checkpoint.com/supportcenter/portal?eventSubmit_doGoviewsolutiondetails=&solutionid=sk30383
3. http://networkgeekstuff.com/networking/introduction-to-checkpoint-firewall-cli-tool-dbedit-and-quick-lab-examples/
4. https://medium.com/@IrekRomaniuk/creating-a-source-for-external-dynamic-list-on-paloalto-firewall-578363f307a8
5. https://www.cpug.org/forums/showthread.php/22008-Split-Tunneling-based-on-Application-Control?p=96108#post96108
6. https://www.cpug.org/forums/showthread.php/21002-How-to-handle-Office365-IP-addresses
