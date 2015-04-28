#!/bin/bash

a="hello lua"

function run {
	echo "hello lua"
}

#run

#for day in Sun Mon Tue Wed Thu Fri Sat
#do
#	echo $day
#done

i=0

max=100
#for  ((i=1;i<=$max;i++))
#do
#	echo $i
#done

#for line in `cat file.txt`  
for line in `cat url.txt`  
do   
	#echo $line | python main.py
	if [ ${#line} -gt 20 ]
	then
		((i++))
		#echo $i
		#echo ${#line}
		#echo $line | python main.py
		#`mv new.mp3 $i.mp3`
		#echo $i.mp3
		#					帐号			bucket			名字		本地路径
		#./qnutil upload -auth=auth.txt -bucket=mua-music -key=$i.mp3 -path=$i.mp3
	fi
	#./main
done
