#! /bin/bash


echo "Siz kiritgan son , minutlarni anglatadi , va shu vaqtdan boshlab ekranni screenshot qilishni boshlaydi , agar to'xtatishni xohlasangiz 0 kiritishingiz mumkin!"

counter=1
t=2
while [ $counter -ne $t ]
do
  read -p "son kiriting:(0 - chiqish) " i 
  if [ $i -eq 0 ]
  then 
     echo "siz dasturni tugatish commandini kiritdingiz"
     counter=2
  else
     sleep $i
     $(gnome-screenshot)
     time="$(date +%F) $(date +%H)-$(date +%M)-$(date +%S)"
     path="/home/x/Pictures/Screenshot from ${time}.png"
     whoami="$(whoami)"
 
     if [ -d "/home/$(whoami)/Documents/backend/Project2/$(date +%F)" ]
     then 
         $(mv "/home/${whoami}/Pictures/Screenshot from ${time}.png"  "/home/$(whoami)/Documents/backend/Project2/screenshot data/$(date +%F)")
     else
     		$(mkdir -p "/home/$(whoami)/Documents/backend/Project2/screenshot data/$(date +%F)")
         $(mv "/home/${whoami}/Pictures/Screenshot from ${time}.png"  "/home/${whoami}/Documents/backend/Project2/screenshot data/$(date +%F)")
     fi

        
  fi
done

#Screenshot from 2022-04-18 21-20-12.png




































