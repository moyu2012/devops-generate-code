#!/usr/bin/expect -f

set timeout -1
set appname [lindex $argv 0]
spawn vue init webpack $appname
expect {
    "($appname)" {
        send "\r";exp_continue
    }
    "(A Vue.js project)" {
        send "\r";exp_continue 
    }
    "Author" {
        send "devops\r";exp_continue
    }
    "Runtime + Compiler: recommended for most users" {
        send "\r";exp_continue
    }
    "(Y/n)" {
        send "n\r";exp_continue
    }
    "NPM" {
        send "y\r";
    }
}

expect eof

exit