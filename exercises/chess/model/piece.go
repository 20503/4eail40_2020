package model

//Implement game pieces here
type piece(X, Y int)(color bool)

type pion(p piece)(alavie, number int)

fonc (p *pion) Mouve(){
  b MouvePiece(X, Y, X, Y++)
  Y++
}
  
type tour(p piece)(alavie, number bool)
fonc (t *tour) Mouve(direction, distence int){//dirction: 1 haut/bas, 2 droite/gauche
  swicth dirction {
    case 1:
      b MouvePiece(X, Y, X, Y+distence)
      Y := Y+distence
    case 2:
      b MouvePiece(X, Y, X+distence, Y)
      X := X+distence
  }
}

type cavalier(p piece)(alavie, number bool)
fonc (c *cavalier) Mouve(pose int){//pose 1a8
  swicth dirction {
    case 1:
      b MouvePiece(X, Y, X+1, Y+2)
      X := X+1
      Y := Y+2
    case 2:
      b MouvePiece(X, Y, X+2, Y+1)
      X := X+2
      Y := Y+1
    case 3:
      b MouvePiece(X, Y, X+2, Y-1)
      X := X+2
      Y := Y-1
    case 4:
      b MouvePiece(X, Y, X+1, Y-2)
      X := X+1
      Y := Y-2
    case 5:
      b MouvePiece(X, Y, X-1, Y-2)
      X := X-1
      Y := Y-2
    case 6:
      b MouvePiece(X, Y, X-2, Y-1)
      X := X-2
      Y := Y-1
    case 7:
      b MouvePiece(X, Y, X-2, Y+1)
      X := X-2
      Y := Y+1
    case 8:
      b MouvePiece(X, Y, X-1, Y+2)
      X := X-1
      Y := Y+2
  }
}

type fou(p piece)(alavie, number bool)
fonc (t *tour) Mouve(direction, distence int){//dirction: 1 haut/gauche, 2 haut/droite, 3 bas/droite, 4 bas/gauche
  swicth dirction {
    case 1:
      b MouvePiece(X, Y, X-distence, Y+distence)
      X := X-distence
      Y := Y+distence
    case 2:
      b MouvePiece(X, Y, X+distence, Y+distence)
      X := X+distence
      Y := Y+distence
    case 3:
      b MouvePiece(X, Y, X+distence, Y-distence)
      X := X+distence
      Y := Y-distence
    case 4:
      b MouvePiece(X, Y, X-distence, Y-distence)
      X := X-distence
      Y := Y-distence
  }
}
type reine(p piece)(alavie bool)
fonc (r *rein) Mouve(direction, distence int){//dirction: 1 haut/bas, 2 droite/gauche
  swicth dirction {
    case 1:
      b MouvePiece(X, Y, X, Y+distence)
      Y := Y+distence
    cecase 2:
      b MouvePiece(X, Y, X+distence, Y+distence)
      X := X+distence
      Y := Y+distence
    case 3:
      b MouvePiece(X, Y, X+distence, Y)
      X := X+distence
    case 4:
      b MouvePiece(X, Y, X+distence, Y-distence)
      X := X+distence
      Y := Y-distence
    case 5:
      b MouvePiece(X, Y, X, Y-distence)
      Y := Y-distence
    case 6:
      b MouvePiece(X, Y, X-distence, Y-distence)
      X := X-distence
      Y := Y-distence
    case 7:
      b MouvePiece(X, Y, X-distence, Y)
      X := X-distence
    case 8:
      b MouvePiece(X, Y, X-distence, Y+distence)
      X := X-distence
      Y := Y+distence
  }
}


type roi(p piece)(alavie bool)
fonc (k *roi) Mouve(direction, distence bool){//dirction: 1 haut/bas, 2 droite/gauche
  swicth dirction {
    case 0:
      b MouvePiece(X, Y, X, Y-1+2*distence)
      Y := Y+distence
    case 1:
      b MouvePiece(X, Y, X-1+2*distence, Y)
      X := X+distence
  }
}
