flipCoin = document.getElementById("flipButton");
flipStatus = document.getElementById("flipStatus");
flipCoin.addEventListener("click", flip);

function flip() {
    let res;
    if(Math.random() < 0.5) {
        flipStatus.textContent = "Heads";
        res = "head";
    }
    else {
        flipStatus.textContent = "Tails";
        res = "tail";
    }
    let data = {
        coin: res
    };
    fetch("/incValues", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            console.log(result)
        });
    }).catch((error) => {
        console.log(error)
    });
}