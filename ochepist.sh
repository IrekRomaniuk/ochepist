#!/bin/bash 
# Source the Check Point profile for library and paths settings
#
export `grep "CPDIR_PATH=" /etc/init.d/firewall1`
[ -f $CPDIR_PATH/tmp/.CPprofile.sh ] || {
    echo "--- Fatal error: cant find CPprofile.sh !!"
    # We are unable to setup essential variables
    echo "Unable to setup needed variables. Stopping." >> $logfile
    exit 2
}
source $CPDIR_PATH/tmp/.CPprofile.sh

# Setup the MDS Environment
#$MDSDIR/scripts/MDSprofile.sh

CMA = "myCMA"
GROUP = "g-ochepist"
mdsenv $CMA
cd /var/scripts/ochepist/
/var/scripts/ochepist/bin/ochepist_lin -g="$GROUP"
if [ $? -eq 0 ]; then
    echo "NOT CHANGED" 
else
    echo "HAS CHANGED" 
    dos2unix results/*
    dbedit -local -ignore_script_failure -globallock -f "results/$GROUP"-dbedit.txt
fi
echo $?
exit
echo "?!"
if [ -f /opt/CPmds-R77/customers/$CMA/CPsuite-R77/fw1/tmp/manage.lock ] 
        then 
                /opt/CPmds-R77/bin/disconnect_client 
        fi
#fwm load policy_name.W gateway_name 2>&1