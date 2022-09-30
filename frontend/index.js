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
            let coinSide = JSON.parse(data).message;
            if(coinSide == "head") {
                headCount.textContent = "This landed on top."
                tailCount.textContent = ""
            } else if(coinSide == "tail") {
                headCount.textContent = ""
                tailCount.textContent = "This landed on top."
            }
        });
    }).catch((error) => {
        console.log(error)
    });
}