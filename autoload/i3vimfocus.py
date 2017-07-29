import subprocess
import vim

def PythonExecSubprocess():
    name = vim.eval('a:name')
    args = vim.eval('a:args')
    name_args = [name] + args
    subprocess.call(name_args)
