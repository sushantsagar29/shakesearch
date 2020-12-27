const Controller = {
    search: (ev) => {
        ev.preventDefault();
        const form = document.getElementById("form");
        const data = Object.fromEntries(new FormData(form));
        const loader = document.getElementById("loader-wrapper");
        loader.style.display = "flex";
        const response = fetch(`/search?q=${data.query}&sensitive=${data.case!==undefined}`).then((response) => {
            response.json().then((results) => {
                Controller.updateTable(results);
            });
        });
    },

    updateTable: (response) => {
        const tableHead = document.getElementById("table-head");
        const tableBody = document.getElementById("table-body");
        const loader = document.getElementById("loader-wrapper");

        tableHead.innerHTML = `<th colspan="2">Number of matches : ${response.count}</th>`;

        let content = "";
        for (let i = 0; i < response.matches.length; i++) {
            content += (`<tr>`);
            content += (`<td>${i + 1}</td>`);
            content += (`<td>${response.matches[i]}</td>`);
            content += (`</tr>`);
        }

        loader.style.display = "none";
        tableBody.innerHTML = content;
    },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
