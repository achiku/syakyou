int ledPin = 12; 
int boardPin = 13;
String str1 = "Str1 is here. ";
String str2 = "Str2 is here. ";

void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  Serial.println("Serial port initialized");
  pinMode(ledPin, OUTPUT);
  pinMode(boardPin, OUTPUT);
}

void loop() {
  // put your main code here, to run repeatedly:
  Serial.println("The Red LED is Blinking");
  digitalWrite(ledPin, HIGH);   // sets the LED on
  delay(1000);                  // waits for a second
  digitalWrite(boardPin, HIGH);

  Serial.println("Hello, from Arduino");
  Serial.println("The Red LED is Blinking");
  digitalWrite(ledPin, LOW);    // sets the LED off
  delay(1000);                  // waits for a second
  digitalWrite(boardPin, LOW);

  Serial.println(str1 + str2);
}
