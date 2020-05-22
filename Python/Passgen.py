import random, string, pyperclip

def randomname(n):
   return ''.join(random.choices(string.ascii_letters + string.digits, k=n))

pyperclip.copy(randomname(16))