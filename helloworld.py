from random import random, seed


print ("HELLO TRY TO GUESS THE NUMBER 1-100")
answer=0
seed (random())
realanswer=int (random ()*100)
guesscount=0
while answer!=realanswer:
    answer=input("guess: ")
    guesscount+=1 
    if int(answer) > realanswer:
        print ("lower")
    elif int(answer) < realanswer:
        print ("higher")
    else: 
        print ("u got it right!!")
        a=("it took you {} guesses")
        print (a.format (guesscount))
        if guesscount > 13 :
            print ("u suck")
        elif guesscount < 5:
            print ("u good")
        break