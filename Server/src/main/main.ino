/*
  Blink

  Turns an LED on for one second, then off for one second, repeatedly.

  Most Arduinos have an on-board LED you can control. On the UNO, MEGA and ZERO
  it is attached to digital pin 13, on MKR1000 on pin 6. LED_BUILTIN is set to
  the correct LED pin independent of which board is used.
  If you want to know what pin the on-board LED is connected to on your Arduino
  model, check the Technical Specs of your board at:
  https://www.arduino.cc/en/Main/Products

  modified 8 May 2014
  by Scott Fitzgerald
  modified 2 Sep 2016
  by Arturo Guadalupi
  modified 8 Sep 2016
  by Colby Newman

  This example code is in the public domain.

  https://www.arduino.cc/en/Tutorial/BuiltInExamples/Blink
*/
#include <ESP8266WiFi.h>
#include <ArduinoWiFiServer.h>
#include "config/arduino_secrets.h"
#include "DataTool.h"

char ssid[] = SECRET_SSID;
char pass[] = SECRET_PASS;

WiFiServer server(80);

// the setup function runs once when you press reset or power the board
void setup() {
  // initialize digital pin LED_BUILTIN as an output.
  pinMode(LED_BUILTIN, OUTPUT);
  Serial.begin(9600);
  Serial.setDebugOutput(true);    //enable the debugging

  //setup the hostname
  WiFi.hostname("DoorSensor");
  //setup the wifi
  WiFi.begin(ssid, pass);
  Serial.print("Connecting");
  while (WiFi.status() != WL_CONNECTED) //wait til connected
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println();

  Serial.print("Connected, IP address: ");
  Serial.println(WiFi.localIP());
  server.begin()
}


// the loop function runs over and over again forever
void loop() {
    // listen for incoming clients
  WiFiClient client = server.available();
  if (client) {
    Serial.println("We got a client, wow!");
    if (client.connected()) {
      Serial.println("Connected to client");
      String * clientData = DataTool.getDataFromClient(client);
      Serial.println("The client wants to tell the server");
      Serial.println(*clientData);
    }

    // close the connection:
    client.stop();
  }  
  delay(5000);                       // wait for 5 seconds before we accept another connection...
  //idk much about networking, but we should try to prevent large requests many many times
}
