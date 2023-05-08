package puppy

import (
	"fmt"

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

func From12() {
	fmt.Println("I am  from versionn 1.2.0")
}

func From13() {
	fmt.Println("I am  from versionn 1.3.0")
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

//https://github.com/Mahadev2023/dog/commit/ba6c3ad54a06f127c749d013b084d31133c688fe
