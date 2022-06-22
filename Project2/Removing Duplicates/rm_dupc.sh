#!/bin/bash

# DIR="/home/x/Pictures"
read -p "Path kiriting:" DIR


for file1 in ${DIR}/*; do
        for file2 in ${DIR}/*; do
                if [[ $file1 != $file2 ]]; then
                        DIFF=`diff "$file1" "$file2" -q`
                        if [ "${DIFF%% *}" != "Files" ]; then
                                echo "Same: $file1 $file2"
                                echo "Remove: $file2"
                                rm "$file1"
                                break
                        fi
                fi
        done
done
echo "Done."

