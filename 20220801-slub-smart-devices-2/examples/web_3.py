from http.server import BaseHTTPRequestHandler, HTTPServer
from gpiozero import LED

led_1 = LED(4)

hostName = ""
serverPort = 8080

class MyServer(BaseHTTPRequestHandler):

    def _redirect(self, path):
        self.send_response(303)
        self.send_header('Content-type', 'text/html')
        self.send_header('Location', path)
        self.end_headers()

    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        html = '''
           <html>
           <body
            style="width:960px; margin: 20px auto;">
           <h1>My Smart Pi</h1>
           <form action="/" method="POST">
               Turn LED_1 :
               <input type="submit" name="LED_1" value="On">
               <input type="submit" name="LED_1" value="Off">
           </form>
           <form action="/" method="POST">
               Turn LED_2 :
               <input type="submit" name="LED_2" value="On">
               <input type="submit" name="LED_2" value="Off">
           </form>
           </body>
           </html>
        '''
        self.wfile.write(html.encode("utf-8"))

    def do_POST(self):

        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length).decode("utf-8")
        print(post_data)

        if post_data == 'LED_1=On':
            led_1.on()
        elif post_data == 'LED_1=Off':
            led_1.off()
        elif post_data == 'LED_2=On':
            led_2.on()
        elif post_data == 'LED_2=Off':
            led_2.off()

        self._redirect('/')

### Main ###

if __name__ == "__main__":
    webServer = HTTPServer((hostName, serverPort), MyServer)
    print("Server started http://%s:%s" % (hostName, serverPort))

    try:
        webServer.serve_forever()
    except KeyboardInterrupt:
        pass

    webServer.server_close()
    print("Server stopped.")
