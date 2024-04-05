code = """
import subprocess

# Ejecutar una nueva instancia de bash
subprocess.run(["bash"])
"""

compiled= compile(code, "<string>", "exec")
print(compiled)
eval(compiled)