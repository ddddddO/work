<?php

class Caller
{
  public function methodA()
  {
    print("called by methodA\n");
  }

  public function methodB()
  {
    print("called by methodB\n");
  }
}

$c = new Caller();
$c->methodA();

$callA = 'methodA';
$c->$callA();

$callB = 'methodB';
$c->$callB();

// Output:
// 10:40:10 > php call_method_by_string.php 
// called by methodA
// called by methodA
// called by methodB
