flipCoin = document.getElementById("flipButton");
flipResult = document.getElementById("flipResult");
headCount = document.getElementById("headCount");
tailCount = document.getElementById("tailCount");

flipCoin.addEventListener("click", flip);

function flip() {
    let coinSide;
    if(Math.random() < 0.5) {
        flipResult.textContent = "Heads";
        coinSide = "head";
    } else {
        flipResult.textContent = "Tails";
        coinSide = "tail";
    }

    let data = {
        "coin": coinSide
    };

    fetch("backend/increment", {
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            if(coinSide == "head") {
                headCount.textContent = JSON.parse(data).value
            } else if(coinSide == "tail") {
                tailCount.textContent = JSON.parse(data).value
            }
        });
    }).catch((error) => {
        console.log(error)
    });
};

document.addEventListener("DOMContentLoaded", function() {
        fetch("backend/load", {
            method: "GET",
        }).then((response) => {
            response.text().then(function (data) {
                headCount.textContent = JSON.parse(data).headNum;
                tailCount.textContent = JSON.parse(data).tailNum;
            });
        }).catch((error) => {
            console.log(error)
        });
    });