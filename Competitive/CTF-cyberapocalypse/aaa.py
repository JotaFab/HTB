import socket
import time

# Función para enviar la respuesta adecuada según el escenario
def send_response(scenarios):
    response = ""
    for scenario in scenarios:
        if scenario == "GORGE":
            response += "STOP-"
        elif scenario == "PHREAK":
            response += "DROP-"
        elif scenario == "FIRE":
            response += "ROLL-"
    # Eliminar el guion al final, si lo hay
    response = response.rstrip("-")
    return response

# Dirección y puerto del servidor del juego
HOST = '94.237.56.248'
PORT = 51227

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    s.sendall(b'y\n')

    # Leer los mensajes del servidor y enviar respuestas
    while True:
        time.sleep(0.1)
        data = s.recv(1024).decode('utf-8')
        old_data=data
        print(data)
        if "What do you do?" in data:
            # Extraer los escenarios de la línea anterior
            
            previous_line = data.split('\n')[-2]
            if ',' in previous_line:
                scenarios = previous_line.split(', ')
            else:
                scenarios = [previous_line]
            response = send_response(scenarios)
            print(response)
            s.sendall(response.encode('utf-8') + b'\n')
        elif "HTB" in data:
            print(data)