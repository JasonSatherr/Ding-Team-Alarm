// DataTool.h
#ifndef DataTool_h
#define DataTool_h

#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <ArduinoWiFiServer.h>
class DataTool {
  private:

  public:
    //DataTool();
    static *getDataFromClient(WiFiClient c);
//    void turnOFF();
//    int getState();
};

#endif
