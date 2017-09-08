
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
$ ls -l results/g-ochepist-dbedit.txt
$ ./ochepist.sh 
1
-rw-rw---- 1 admin root 423726 Sep  6 13:34 results/g-ochepist-dbedit.txt
$ ./ochepist.sh 
0
$ ls -l results/g-ochepist-dbedit.txt
-rw-rw---- 1 admin root 423726 Sep  6 13:34 results/g-ochepist-dbedit.txt
$./ochepist
....
g-ochepist updated successfully.
Database saved successfully
161
### Example

```
$ ./ochepist_linux_386 -url="https://minemeld/feeds/office365_IPv4s" -g="g-o365"

$ dbedit -local -ignore_script_failure -globallock -f results/g-o365-dbedit.txt

```
### ToDo

+ add #comments in template files i.e. do not add spaces or empty lines

### References

1. https://sc1.checkpoint.com/documents/R77/CP_R77_CLI_ReferenceGuide_WebAdmin/html_frameset.htm?topic=documents/R77/CP_R77_CLI_ReferenceGuide_WebAdmin/105997
2. https://supportcenter.checkpoint.com/supportcenter/portal?eventSubmit_doGoviewsolutiondetails=&solutionid=sk30383
3. http://networkgeekstuff.com/networking/introduction-to-checkpoint-firewall-cli-tool-dbedit-and-quick-lab-examples/
4. https://medium.com/@IrekRomaniuk/creating-a-source-for-external-dynamic-list-on-paloalto-firewall-578363f307a8
5. https://www.cpug.org/forums/showthread.php/22008-Split-Tunneling-based-on-Application-Control?p=96108#post96108
6. https://www.cpug.org/forums/showthread.php/21002-How-to-handle-Office365-IP-addresses
