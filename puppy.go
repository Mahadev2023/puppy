package puppy

import (
	"github.com/Mahadev2023/dog"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return `Woof Woof ...`
}

func BigBarks() string {
	return dog.WhenGrounUp(Bark())
}

// https://github.com/Mahadev2023/puppy
//eval `ssh-agent -s`

//eval "$(ssh-agent -s)"

//to set github
/*
ssh-keygen -t ed25519 -C "godbolemahadevkanaka@gmail.com"
Generating public/private ed25519 key pair.
Enter file in which to save the key (/c/Users/godbo/.ssh/id_ed25519):
*/

//to add the file in git do this always
//eval "$(ssh-agent -s)"
//ssh-add ~/.ssh/id_ed25519
//
