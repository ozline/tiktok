
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=api
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
