<?php

class Client {

    private $conn;

    function __construct($host, $port) {
        $this->conn = fsockopen($host, $port, $errno, $php_errormsg, 3);
        if ($this->conn) {
            return false;
        }
    }

    public function Call($method, $params) {
        if (!$this->conn) {
            return false;
        }

        $err = fwrite($this->conn, json_encode(array(
            "method" => $method,
            "params" => array($params),
            "id" => 0,
        ))."\n");

        if ($err === false) {
            return false;
        }

        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);

        if ($line === false) {
            return false;
        }

        return json_decode($line, true);
    }
}

$client = new Client("127.0.0.1", "8096");
$args = array("A" => 9, "B" => 2);
$res = $client->Call("Rect.Area", $args);
print_r($res["result"]);
printf("%d * %d = %d\n", $args["A"], $args["B"], $res["result"]);