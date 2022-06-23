// LED.cpp
#include "LED.h"
/* duh constructor of old
LED::LED(int pin) {
  ledPin   = pin;
  ledState = LOW;
  pinMode(ledPin, OUTPUT);
}
*/
static * DataTool::getDataFromClient(client c) {
  String dataString = String("");
  nextChar = c.read()
  while (nextChar != -1){
    dataString += nextChar; //append the next char
    nextChar = c.read();    //read the next char from the client
  }
  
}
/*
void LED::turnOFF() {
  ledState = LOW;
  digitalWrite(ledPin, ledState);
}

int LED::getState() {
  return ledState;
}
*/
