<?php
    if ($_SERVER["REQUEST_METHOD"] != "GET" || !isset($_GET["filename"])) {
        return;
    }
    $path = pathinfo($_GET["filename"]);

    if (!rename($path["filename"].".txt",$path["filename"].".lua")) {
        echo "Could not set the file name";
        return;
    }

    echo "Changed the file name";
?>