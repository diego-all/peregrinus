import platform
import os
import subprocess

def get_system_info():
    info = {}

    # Obtener la versión del sistema operativo
    info['OS Version'] = platform.platform()

    # Obtener la arquitectura del sistema
    info['Architecture'] = platform.architecture()[0]

    # Obtener usuarios del sistema
    try:
        users = subprocess.check_output("cut -d: -f1 /etc/passwd", shell=True).decode().splitlines()
        info['Users'] = users
    except Exception as e:
        info['Users'] = f"Error al obtener usuarios: {str(e)}"

    # Obtener programas instalados (en sistemas basados en Debian/Ubuntu)
    try:
        installed_programs = subprocess.check_output("dpkg --list", shell=True).decode()
        info['Installed Programs'] = installed_programs
    except Exception as e:
        info['Installed Programs'] = f"Error al obtener programas instalados: {str(e)}"

    return info

def main():
    info = get_system_info()
    for key, value in info.items():
        print(f"{key}:")
        if isinstance(value, list):
            for item in value:
                print(f"  {item}")
        else:
            print(f"  {value}")

if __name__ == "__main__":
    main()
