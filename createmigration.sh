#!/bin/sh

while getopts :d: name
do
    case $name in
    d)    description="$OPTARG";;
  	?)    printf "Usage: %s [-d \"migration description\"]\n" $0
  		    exit 2;;
    esac
done

if [ -z "$description" ]; then
    echo "Error: description must not be empty"
    exit
fi

location="$PWD/database/migration/"
version="$(date +%Y%m%d%H%M%S)"
suffix=".go"

filename="${version}__$( printf "$description" | sed 's/ /_/g' )$suffix"

echo "package migration\n\nimport (\n\t\"github.com/go-gormigrate/gormigrate/v2\"\n)\n\nvar V${version} = gormigrate.Migration{\n\tID: \"V${version}\",\n}" > $location$filename
echo "Migration successfully created"
echo "Migration file: file://$location$filename"