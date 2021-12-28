package main

func myFunction(ch chan int) { } // passing a bidirectorional channel

func myOtherfunction(ch chan<- int){} // passing channel as send-only

func myReceiverFunction(ch <-chan int) { } // receive-only channel