import sys, pygame
pygame.init()
pygame.font.init()

size = width, height = 800, 600
black = 0, 0, 0
x = 200
y = 200

screen = pygame.display.set_mode(size)



while 1:
    pygame.time.delay(10)
    for event in pygame.event.get():
        if event.type == pygame.QUIT: sys.exit()
    keys = pygame.key.get_pressed() 
    if keys[pygame.K_LEFT] and x>0: 
        x -= 10
    if keys[pygame.K_RIGHT] and x<width: 
        x += 10
    if keys[pygame.K_UP] and y>0: 
        y -= 10
    if keys[pygame.K_DOWN] and y<height: 
        y += 10

        print(x)

    screen.fill(black)
    pygame.draw.circle(screen, (200, 0, 0), (x, y), 80)
    pygame.display.flip()