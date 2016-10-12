// +build !arm
//
// gpio mock
//  - to deploy under x86
//  - the real implementation depends on the bcm2835 library

package gpio

import (
	"log"
	"os"
	"os/signal"
	"time"
)

/*
#include <sys/ioctl.h>
#include <sys/select.h>
#include <stdio.h>
#include <unistd.h>
#include <termios.h>

void echo(){
   struct termios old = {0};
   fflush(stdout);
   if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
   old.c_lflag |= ICANON;
   old.c_lflag |= ECHO;
   if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
}

void noecho(){
   struct termios old = {0};
   fflush(stdout);
   if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
   old.c_lflag &= ~ICANON;
   old.c_lflag &= ~ECHO;
   old.c_cc[VMIN] = 1;
   old.c_cc[VTIME] = 0;
   if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
}

char getch(){
   char ch = 0;
   struct termios old = {0};
   fflush(stdout);
   if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
   old.c_lflag &= ~ICANON;
   old.c_lflag &= ~ECHO;
   old.c_cc[VMIN] = 1;
   old.c_cc[VTIME] = 0;
   if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
   if( read(0, &ch,1) < 0 ) perror("read()");
   old.c_lflag |= ICANON;
   old.c_lflag |= ECHO;
   if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
   return ch;
}

int kbhit() {
  struct timeval tv;
  fd_set fds;
  tv.tv_sec = 0;
  tv.tv_usec = 0;
  FD_ZERO(&fds);
  FD_SET(STDIN_FILENO, &fds); //STDIN_FILENO is 0
  select(STDIN_FILENO+1, &fds, NULL, NULL, &tv);
  return FD_ISSET(STDIN_FILENO, &fds);
}

*/
import "C"

func init() {
	// disable console echo
	go func() {
		// HACK (this is only for dev / test, so i'm fine with this here)
		// wait 1 second before disable echo
		// (do prevent disable console echo if the app exits after startup (-h flag))
		time.Sleep(time.Second * 1)
		C.noecho()
	}()

	// re-enable console echo on exit (Ctrl-C)
	e := make(chan os.Signal, 1)
	signal.Notify(e, os.Interrupt)
	go func() {
		for _ = range e {
			C.echo()
			os.Exit(0)
		}
	}()
}

func Read(pin Pin) (LowOrHigh, error) {
	value := Low
	if C.kbhit() != 0 {
		C.getchar()
		value = High
	}

	log.Printf("Read(pin: %d) - return value: '%s'", pin, Low)
	return value, nil
}

func Write(pin Pin, value LowOrHigh) error {
	log.Printf("Write(pin: %d, value: %s)\n", pin, value)
	return nil
}

func Pwm(pin Pin, value int) error {
	log.Printf("Pwm(pin: %d, value: %d)\n", pin, value)
	return nil
}

func BlockWhileIsLow(pin Pin) (bool, time.Duration, error) {
	start := time.Now()
	isLow := false

	for C.kbhit() != 0 {
		isLow = true
		C.getchar()
		time.Sleep(time.Millisecond * 50)
	}

	duration := time.Since(start)

	if isLow {
		log.Printf("IsLow(isLow: %t, duration: %s\n", isLow, duration)
	}

	return isLow, duration, nil
}
