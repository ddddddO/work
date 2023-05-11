#!/usr/bin/env sh

# set -u


header() {
  echo 'ID,Name,Address'
}

# 引数：$1=start, $2=range, $3=USE_ID, $4=USE_ADRRESS, $5=APPEND_NAME_STR, $6=APPEND_ADRRESS_STR
records() {
  start=$1
  end=`expr $start + $2 - 1`
  use_id=$3
  use_address=$4
  append_name_str=$5
  append_address_str=$6

  for i in `seq $start $end`
  do
    id=""
    if [ $use_id = "true" ]; then
      id="${i}"
    fi
    address=""
    if [ $use_address = "true" ]; then
      address="address_${i}${append_address_str}"
    fi
    name="name_${i}${append_name_str}"

    echo "${id},${name},${address}"
  done
}

usage() {
  cat 1>&2 <<EOF
Usage: $(basename $0) [OPTIONS]

Options:
  -i        Generate ID
  -a        Generate Address

  -x        Append specified string to Name
  -z        Append specified string to Address

  -s        Specify start number
  -r        Specify range from the start number
              e.g.) $ ./csv.sh -s 3 -r 5
                      => 3,4,5,6,7
  -h        Help
EOF
  exit 1
}


USE_ID=false
USE_ADRRESS=false
APPEND_NAME_STR=""
APPEND_ADRRESS_STR=""
START=0
RANGE=0

# 「引数ありのオプションにしたいときは、オプション文字列の対象文字の後ろに : を付けます。」
while getopts 'aix:z:s:r:h' flag; do
  case "${flag}" in
    i) USE_ID=true ;;
    a) USE_ADRRESS=true ;;

    x) APPEND_NAME_STR=$OPTARG ;;
    z) APPEND_ADRRESS_STR=$OPTARG ;;

    s) START=$OPTARG ;;
    r) RANGE=$OPTARG ;;

    h) usage ;;
    ?) usage ;;
  esac
done

header
records $START $RANGE $USE_ID $USE_ADRRESS $APPEND_NAME_STR $APPEND_ADRRESS_STR
