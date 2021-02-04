<?php
 /**
  * @var Goridge\RelayInterface $relay
  */

 use Spiral\Goridge;
 use Spiral\RoadRunner;

 $rr = new RoadRunner\Worker($relay);

 while ($in = $rr->waitPayload()) {
     try {
         $rr->send((string)getmypid());
     } catch (\Throwable $e) {
         $rr->error((string)$e);
     }
 }
