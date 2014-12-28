button
======

Go language support of Dream Cheeky's Big Read Button (http://dreamcheeky.com/big-red-button).

You need a udev rule to access this device.  Add this file:
/etc/udev/rules.d/50-big-red-button.rules

It should contain:
ACTION=="add", ENV{ID_MODEL}=="DL100B_Dream_Cheeky_Generic_Controller", SYMLINK+="big_red_button", MODE="0666"
ACTION=="remove", ENV{ID_MODEL}=="DL100B_Dream_Cheeky_Generic_Controller"


I learned how to do this from this page:
http://blog.opensensors.io/blog/2013/11/25/the-big-red-button/

I'm heavily tweaking this file so it may be in any state.  I'm using Github primarily for backup
but if this helps you then great!

