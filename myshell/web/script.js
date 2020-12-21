const commandText = document.getElementById("command")
const sendCommandButton = document.getElementById("sendCommand");
const outputParagraph = document.getElementById('output');

sendCommandButton.addEventListener("click", function () {
    let command = commandText.value;
    let url = 'http://localhost:8080/' + command;
    fetch(url, { method: "POST", body: command })
        .then(response => response.text())
        .then((response) => {
            console.log(response);
            // outputParagraph.innerHtml = '<span class="result">' + response + '</span>' + outputParagraph.innerText;
            // outputParagraph.innerText = '<span class="command">' + command + '</span>' + outputParagraph.innerText;
            outputParagraph.innerHTML = '<p class="command">' + command + '</p>' + '<p class="result">' + response + '</p>' + outputParagraph.innerHTML;
        })
    .catch(err => console.log(err))

});