/*
 * This example assumes the program you want to execute
 * is called child.exe, and resides in the same 
 * directory as this program
 */

#include <stdlib.h> 
#include <stdio.h> 
#include <string.h> 

int main()
{

  char child2[BUFSIZ];

  strcpy (child2, "/home/dssconfirm/C_test/dssconfirm");
  strcat (child2, " -m iccid");

  printf ("Executing %s\n", child2);
  int ret = system (child2);
  printf ("Return code: %d\n", ret);
  return 0;
}
