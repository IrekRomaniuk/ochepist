
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

### Deploy on Linux

```
$ mkdir -p ochepist                
$ cd ochepist
$ tar zxvf ochepist.tar.gz         
bin/
bin/ochepist_lin
bin/ochepist_win
templates/
templates/hosts.gotxt
templates/networks.gotxt
templates/ranges.gotxt
results/g-ochepist-test-dbedit.txt

$ chmod +x bin/ochepist_lin 
$ bin/ochepist_lin -url="O365IPAddresses.xml" -g="g-o365"

$ head results/g-o365-dbedit.txt 
delete network_objects g-o365
create network_object_group g-o365
delete network_objects n13.64.196.27-255.255.255.255
create network n13.64.196.27-255.255.255.255
modify network_objects n13.64.196.27-255.255.255.255 ipaddr 13.64.196.27
modify network_objects n13.64.196.27-255.255.255.255 netmask 255.255.255.255
modify network_objects n13.64.196.27-255.255.255.255 comments "Created by ochepist with dbedit"
update network_objects n13.64.196.27-255.255.255.255
addelement network_objects g-o365 '' network_objects:n13.64.196.27-255.255.255.255
update network_objects g-o365

$ cat results/g-o365-dbedit.txt | wc -l
1411


```

### Example

```
Reads addresses from url and writes in results/g-o365-dbedit.txt. To read from https://support.content.office.net/en-us/static/O365IPAddresses.xml use 'O365IPAddresses.xml' as url

$ ./ochepist_linux_386 -url="https://minemeld/feeds/office365_IPv4s" -g="g-o365"

Imports into Checkpoint management server

$ dbedit -local -ignore_script_failure -globallock -f results/g-o365-dbedit.txt

```
### ToDo

+ add #comments in template files i.e. do not add spaces or empty lines

### References
1. Split tunneling based on external dynamic lists for Checkpoint https://medium.com/@IrekRomaniuk/split-tunneling-based-on-external-dynamic-lists-for-checkpoint-2ec20aab6525
2. https://sc1.checkpoint.com/documents/R77/CP_R77_CLI_ReferenceGuide_WebAdmin/html_frameset.htm?topic=documents/R77/CP_R77_CLI_ReferenceGuide_WebAdmin/105997
3. https://supportcenter.checkpoint.com/supportcenter/portal?eventSubmit_doGoviewsolutiondetails=&solutionid=sk30383
4. http://networkgeekstuff.com/networking/introduction-to-checkpoint-firewall-cli-tool-dbedit-and-quick-lab-examples/
4. https://medium.com/@IrekRomaniuk/creating-a-source-for-external-dynamic-list-on-paloalto-firewall-578363f307a8
6. https://www.cpug.org/forums/showthread.php/22008-Split-Tunneling-based-on-Application-Control?p=96108#post96108
7. https://www.cpug.org/forums/showthread.php/21002-How-to-handle-Office365-IP-addresses
