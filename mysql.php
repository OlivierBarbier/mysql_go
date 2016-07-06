<?php

$dsn = 'mysql:dbname=tiller;host=127.0.0.1';
$user = 'root';
$password = 'root';

try {
    $dbh = new PDO($dsn, $user, $password);

    echo "We are good to go ;)";

    $stmt = $dbh->prepare("INSERT INTO squareNum (val) VALUES(?) ");

    for($i = 0; $i < 5000; $i++) {
        $stmt->execute([1]);
    }
} catch (PDOException $e) {
    echo 'Connection failed: ' . $e->getMessage();
}
echo ". Done!\n";
