#!/bin/sh

set -e
ls -lh '/opt/my_backup.sql'
#pg_restore --verbose --clean --create -U thder77777 -d database1 '/opt/my_backup.sql'

psql -d database1 -f '/opt/my_backup.sql'
