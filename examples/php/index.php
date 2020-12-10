<?php

use MuenchDev\GrpcDemo\StockUpdateRequest;

require 'vendor/autoload.php';

$client = new MuenchDev\GrpcDemo\StockClient(
    'host.docker.internal:9000',
    ['credentials' => Grpc\ChannelCredentials::createInsecure()],
);

$unaryCall = $client->Update(
    (new StockUpdateRequest())
        ->setQty((float) mt_rand(0, 1000))
        ->setSku('24-MB01')
        ->setSourceCode('default')
);

echo "<pre>";
echo sprintf("Peer: %s\n\n", $unaryCall->getPeer());

$response = $unaryCall->wait();
echo sprintf("ACK: %b\n", $response[0]->getAck());

echo "</pre>";
