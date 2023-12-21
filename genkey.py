import random, string

def genkey():
    lcnkey = ""
    for anu in range(4):
        block = ''.join(random.choices(string.ascii_uppercase + string.digits, k=4))
        lcnkey += block + "-"
    lcnkey = lcnkey[:-1]
    return lcnkey

print(genkey())
