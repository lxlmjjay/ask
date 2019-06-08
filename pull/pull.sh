#!/bin/bash
cd /root/go/src/htdocs/ask && git pull

if [ $? -ne 0 ]; then
	echo "failed"
else
	echo "succeed"
fi
