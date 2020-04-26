# live-check
It just responds 200 OK to :8080 / . Literally nothing else

## Install as a service
Run these commands to get livecheck running as a service. These assume that you're in the
directory that contains both the `livecheck` binary and the `livecheck.service` unit file and that you are running these commands as root.

```Shell
chown root:root livecheck*
chmod 555 livecheck
chmod 755 livecheck.service
mkdir -p /opt/livecheck/bin
mv livecheck /opt/livecheck/bin
mv livecheck.service /lib/systemd/system/

systemctl enable livecheck.service
systemctl start livecheck
```