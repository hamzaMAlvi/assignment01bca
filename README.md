# assignment01bca
Solution of Assignment 1 in the Course of Blockchain and its Applications offered in Spring 2020 at FAST-NU.

The Specifics given for the assignment are:

Your submission should be made on the classroom and you should submit a zip file containing both the main and assignment01bca package (see below). Deadline is 17th February 12:55pm. Late submissions or ones sent by email are NOT accepted and WILL give you zero points. You should use the following naming convention while uploading file:

assignment01bca_YourRollNO.zip

Before submission, you need to create a golang package named exactly as assignment01bca that should be available on the Github. It should have at least following public functions:

func InsertBlock(transaction string, chainHead *block) *block {
func ListBlocks(chainHead *block) {
func ChangeBlock(oldTrans string, newTrans string, chainHead *block) {
func VerifyChain(chainHead *block) {

The following code, when your github username is placed at <username>, should work and display all the info to evaluate the assignment. Please note that you risk losing ALL your marks if the code below doesnot work for the package you submit. You should make your github repo private and add me (ehteshamz) as a collaborator.

package main

import (
a1 "github.com/<username>/assignment01bca"
)

func main() {
  var chainHead *a1.Block
  chainHead = a1.InsertBlock("GenesisBlock", nil)
  chainHead = a1.InsertBlock("AliceToBob", chainHead)
  chainHead = a1.InsertBlock("BobToCharlie", chainHead)
  a1.ListBlocks(chainHead)
  a1.ChangeBlock("AliceToBob", "AliceToTrudy", chainHead)
  a1.ListBlocks(chainHead)
  a1.VerifyChain(chainHead)
}
