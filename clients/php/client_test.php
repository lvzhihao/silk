<?php

require dirname(__FILE__) . "/vendor/autoload.php";

@include_once dirname(__FILE__) . "/Silk/Core/AccountRequest.php";
@include_once dirname(__FILE__) . "/Silk/Core/AccountResponse.php";
@include_once dirname(__FILE__) . "/Silk/Core/SilkerClient.php";
@include_once dirname(__FILE__) . "/GPBMetadata/Accounts.php";

use Google\Protobuf\Internal\MapField;
use Google\Protobuf\Internal\GPBType;

$client = new Silk\Core\SilkerClient("localhost:8588", [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

$setValue = function($var) {
   return $this->setValue($var);
};

$anyPack = function($var) {
    $this->pack($var);
    return $this;
};

$request = new Silk\Core\AccountRequest();
$request->setPlatform("uchat");
$request->setMetadata(json_encode([
    "key1" => "string",
]));
list($response, $status) = $client->CreateAccount($request)->wait();
if($status->code > 0) {
    var_dump($status);
} else {
    $id = $response->getId();
    echo $id, "\n";
}
